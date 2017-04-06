// Comparing memory usage between a struct vs a pointer in a channel

package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"runtime/debug"
	"strconv"
	"time"
	// "unsafe"

	"github.com/dustin/go-humanize"
	"github.com/pkg/profile"
)

var _ = fmt.Sprint
var _ = time.Now()
var _ = debug.FreeOSMemory

type SmallStruct struct {
	Array [25]byte
}

type MediumStruct struct {
	Array [100]byte
}

type LargeStruct struct {
	Array [500]byte
}

func printMemStats(stats runtime.MemStats) {
	fmt.Printf("-- MemStats --\n%10s: %s\n%10s: %s\n%10s: %s\n%10s: %s\n%10s: %s\n%10s: %s\n%10s: %s\n",
		"HeapSys", humanize.Bytes(stats.HeapSys),
		"HeapInuse", humanize.Bytes(stats.HeapInuse),
		"Alloc", humanize.Bytes(stats.Alloc),
		"HeapAlloc", humanize.Bytes(stats.HeapAlloc),
		"HeapIdle", humanize.Bytes(stats.HeapIdle),
		"HeapObjs", humanize.Bytes(stats.HeapObjects),
		"HeapRel", humanize.Bytes(stats.HeapReleased),
	)
}

var printOut = false

func checkMemstats() {
	for {
		if printOut {
			var mem runtime.MemStats
			runtime.ReadMemStats(&mem)
			printMemStats(mem)
		}
		time.Sleep(1 * time.Second)
	}
}

func control() {
	fmt.Println("Starting Control")
	for true {
		in, err := bufio.NewReader(os.Stdin).ReadBytes('\n')
		if err != nil {
			fmt.Println(err)
		}
		cmd := string(in)[:len(in)-1]

		switch {
		case cmd == "m":
			printOut = !printOut
			fmt.Println(printOut)
		case len(cmd) > 1 && cmd[0] == 'd':
			i, err := strconv.Atoi(cmd[1:])
			if err == nil {
				_ = make(chan SmallStruct, 100000*i)
				fmt.Printf("Made %s bytes\n", humanize.Bytes(uint64(i*100000*25)))
			} else {
				fmt.Println(err)
			}
		case cmd == "q" || cmd == "quit" || cmd == "exit":
			return
		case cmd == "r":
			debug.FreeOSMemory()
			fmt.Println("Free")
		case cmd == "g":
			runtime.GC()
			fmt.Println("GC")
		default:
			fmt.Println("Command not found")
		}
	}
}

func main() {
	go checkMemstats()
	runtime.MemProfileRate = 1
	p := profile.Start(profile.MemProfileRate(1), profile.ProfilePath("."))
	defer p.Stop()

	/*fmt.Printf("SmallSize: %d\nMediumSize:%d\nLargeSize:%d\n", int(unsafe.Sizeof(SmallStruct{})),
	int(unsafe.Sizeof(MediumStruct{})),
	int(unsafe.Sizeof(LargeStruct{})))*/

	sa := make(chan SmallStruct, 100000)
	sb := make(chan *SmallStruct, 100000)
	sc := make(chan SmallStruct, 1)
	sd := make(chan *SmallStruct, 1)

	ma := make(chan MediumStruct, 100000)
	mb := make(chan *MediumStruct, 100000)
	mc := make(chan MediumStruct, 1)
	md := make(chan *MediumStruct, 1)

	la := make(chan LargeStruct, 100000)
	lb := make(chan *LargeStruct, 100000)
	lc := make(chan LargeStruct, 1)
	ld := make(chan *LargeStruct, 1)

	control()

	//debug.FreeOSMemory()
	runtime.GC()

	fmt.Println(sa, sb, sc, sd,
		ma, mb, mc, md,
		la, lb, lc, ld,
	)
}
