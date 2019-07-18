package benchmarks_test

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/pprof"
	_ "net/http/pprof"
	"os"
	"runtime"
	"testing"
	"time"

	"golang.org/x/time/rate"

	"github.com/gonum/floats"
	"github.com/struCoder/pidusage"

	"github.com/FactomProject/factomd/common/entryBlock"
	"github.com/FactomProject/factomd/common/primitives"
	"github.com/FactomProject/factomd/common/primitives/random"
	linuxproc "github.com/c9s/goprocinfo/linux"
)

// go test jsonencoding_test.go -bench=.

const (
	Byte = 1
	Kb   = 2 ^ 10*Byte
	Mb   = 2 ^ 10*Kb
)

// go test -run TestCPUProfile

func TestCPUProfile(t *testing.T) {
	runtime.GOMAXPROCS(1)
	marshalled := 0
	quit := make(chan int)
	go StartProfiler(16060)
	go func() {
		lastM := 0
		start := time.Now()
		tickerDur := time.Second
		ticker := time.NewTicker(tickerDur)
		p := GetStat()
		for tick := range ticker.C {
			m := marshalled
			elapsed := tick.Sub(start)

			fmt.Printf("Elapsed %s. Marshaled %d entries. Total %d, Avg %.2f/s\n", elapsed, m-lastM, m, float64(m)/elapsed.Seconds())
			lastM = m
			c := GetStat()
			ReportCPUUtilization(p, c, tickerDur)
			p = c

			if elapsed > time.Second*20 {
				// Stop the test
				fmt.Println()
				fmt.Printf("Over the Entire Run:\n")
				for k, avg := range Averages {
					fmt.Printf("\t%s: %.2f%%\n", k, floats.Sum(avg)/float64(len(avg)))
				}
				quit <- 0
				return
			}
		}
	}()
	// First construct a set of random entries
	ents := RandomEntries(1000, Kb, 5*Kb, 5, 50*Byte)

	// Do a throughput of X/min
	x := 5000
	RunMode := 7
	rateLimiter := rate.NewLimiter(rate.Limit(x), 10)
	ctx := context.Background()
	c := 0

	jsonEncode := func(e *entryBlock.Entry) ([]byte, error) {
		return json.Marshal(e)
	}
	marshalEncode := func(e *entryBlock.Entry) ([]byte, error) {
		return e.MarshalBinary()
	}

	encode := jsonEncode

	switch RunMode {
	case 0:
		fmt.Printf("Running the BusyLoop at %d entries per second with JSON encode\n", x)
		goto BusyLoop
	case 1:
		x = 1500
		rateLimiter = rate.NewLimiter(rate.Limit(x), 10)
		fmt.Printf("Running the BusyLoop at %d entries per second with JSON encode\n", x)
		goto BusyLoop
	case 2:
		x = 500
		rateLimiter = rate.NewLimiter(rate.Limit(x), 10)
		fmt.Printf("Running the BusyLoop at %d entries per second with JSON encode\n", x)
		goto BusyLoop
	case 3:
		fmt.Printf("Running the TightLoop at unlimited entries per second with JSON encode\n")
		goto TightLoop
	case 4:
		fmt.Printf("Running the CPUHog in a tight for loop\n")
		goto CPUHog
	case 5:
		fmt.Printf("Running Nada to get a control\n")
		goto Nada
	case 6:
		encode = marshalEncode
		fmt.Printf("Running the TightLoop at unlimited entries per second with Marshal encode\n")
		goto TightLoop
	case 7:
		x = 1500
		encode = marshalEncode
		rateLimiter = rate.NewLimiter(rate.Limit(x), 10)
		fmt.Printf("Running the BusyLoop at %d entries per second with Marshal Encode\n", x)
		goto BusyLoop
	}

BusyLoop:
	for {
		//fmt.Print(".")
		err := rateLimiter.WaitN(ctx, 10)
		if err != nil {
			panic(err)
		}
		for i := 0; i < 10; i++ {
			data, err := encode(ents[c%len(ents)])
			if err != nil {
				panic(err)
			}
			var _ = data
			c++
			marshalled++
		}

		select {
		case <-quit:
			break BusyLoop
		default:

		}
	}
	goto End

TightLoop:
	for {
		data, err := encode(ents[c%len(ents)])
		if err != nil {
			panic(err)
		}
		var _ = data
		c++
		marshalled++

		select {
		case <-quit:
			break TightLoop
		default:
		}
	}
	goto End

CPUHog:
	for {
		select {
		case <-quit:
			break CPUHog
		default:
		}
	}
	goto End

Nada:
	for {
		time.Sleep(1 * time.Second)
		select {
		case <-quit:
			break Nada
		default:
		}
	}

End:
}

var Averages = make(map[string][]float64)

func ReportCPUUtilization(prior, current *linuxproc.Stat, d time.Duration) {
	UserTotal := uint64(0)
	SystemTotal := uint64(0)
	for i, c := range current.CPUStats {
		p := prior.CPUStats[i]
		// Plus 1 so cores line up with monitor
		UserTotal += c.User - p.User
		SystemTotal += c.System - p.System
		fmt.Printf("\t[%d] User %d%%, System %d, Nice %d\n", i+1, c.User-p.User, c.System-p.System, c.Nice-p.Nice)
		// s.User
		// s.Nice
		// s.System
		// s.Idle
		// s.IOWait
	}

	uAvg := float64(UserTotal) / float64(len(current.CPUStats))
	sAvg := float64(SystemTotal) / float64(len(current.CPUStats))
	fmt.Printf("Average of User %.2f%%, System %.2f%%\n",
		uAvg,
		sAvg)

	Averages["user"] = append(Averages["user"], uAvg)
	Averages["sys"] = append(Averages["sys"], sAvg)

	sysInfo, err := pidusage.GetStat(os.Getpid())
	if err != nil {
		panic(err)
	}
	fmt.Printf("PidUsage: %.2f\n", sysInfo.CPU)
	Averages["pid"] = append(Averages["pid"], sysInfo.CPU)
}

func GetStat() *linuxproc.Stat {
	stat, err := linuxproc.ReadStat(fmt.Sprintf("/proc/stat"))
	if err != nil {
		panic("stat read fail")
	}
	return stat
}

func BenchmarkEncodeEntryBinary(b *testing.B) {
	// First construct a set of random entries
	ents := RandomEntries(1000, Kb, 5*Kb, 5, 50*Byte)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		data, err := ents[i%len(ents)].MarshalBinary()
		if err != nil {
			panic(err)
		}
		var _ = data
	}
}

func BenchmarkEncodeEntryJSON(b *testing.B) {
	// First construct a set of random entries
	ents := RandomEntries(1000, Kb, 5*Kb, 5, 50*Byte)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		data, err := json.Marshal(ents[i%len(ents)])
		if err != nil {
			panic(err)
		}
		var _ = data
	}
}

func RandomEntries(amt, minSize, maxSize, extCount, extSize int) []*entryBlock.Entry {
	arr := make([]*entryBlock.Entry, amt)
	for i, _ := range arr {
		arr[i] = RandomEntry(minSize, maxSize, extCount, extSize)
	}
	return arr
}

func RandomEntry(minSize, maxSize, extCount, extSize int) *entryBlock.Entry {
	e := entryBlock.NewEntry()
	e.Version = random.RandUInt8()
	e.ChainID = primitives.RandomHash()
	for i := 0; i < extCount; i++ {
		e.ExtIDs = append(e.ExtIDs, primitives.ByteSlice{Bytes: random.RandByteSliceOfLen(extSize)})
	}

	l := random.RandIntBetween(minSize, maxSize)
	e.Content = primitives.ByteSlice{Bytes: random.RandByteSliceOfLen(l)}
	return e
}

func StartProfiler(logPort int) {
	mux := http.NewServeMux()
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)

	log.Println(http.ListenAndServe(fmt.Sprintf("localhost:%d", logPort), mux))
	//runtime.SetBlockProfileRate(100000)
}
