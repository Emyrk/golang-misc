package benchmarks_test

import (
	"runtime"
	"testing"
	"time"

	state "github.com/Emyrk/golang-misc/benchmarks"
	"github.com/FactomProject/factomd/common/primitives/random"
)

var (
	zero = [32]byte{}
)

var qt = 0

func BenchmarkBalanceMap(b *testing.B) {
	b.Run("Balance Map -- Control (Direct)", benchmarkBalanceMapDirect)

	qt = -1
	b.Run("Balance Map Mutex -- Empty Map Gets", benchmarkBalanceMapEmptyGets)
	b.Run("Balance Map Mutex -- Competing Gets", benchmarkBalanceMapEmptyGetsWithCompeting)

	qt = 0
	b.Run("Balance Map Shared Obj -- Empty Map Gets", benchmarkBalanceMapEmptyGets)
	b.Run("Balance Map Shared Obj -- Competing Gets", benchmarkBalanceMapEmptyGetsWithCompeting)

	qt = 1
	b.Run("Balance Map Chan Back -- Empty Map Gets", benchmarkBalanceMapEmptyGets)
	b.Run("Balance Map Chan Back -- Competing Gets", benchmarkBalanceMapEmptyGetsWithCompeting)

	qt = 2
	b.Run("Balance Map Pooled Chan Back -- Empty Map Gets", benchmarkBalanceMapEmptyGets)
	b.Run("Balance Map Pooled Chan Back -- Competing Gets", benchmarkBalanceMapEmptyGetsWithCompeting)

}

// benchmarkBalanceMapEmptyGetsWithCompeting tests against an empty map, focusing on the channel + query speeds
//	Also has competing threads
func benchmarkBalanceMapEmptyGetsWithCompeting(b *testing.B) {
	n := runtime.NumGoroutine()
	bm := state.NewBalanceMap()
	bm.SetDirectBalance(zero, 10)
	go bm.Serve()
	for i := 0; i < 100; i++ {
		go func() {
			for {
				defer func() {
					if r := recover(); r != nil {
					}
				}()
				if bm.Closed() {
					return
				}
				if bal := bm.GetBalance(zero, qt); bal != 10 {
					b.Errorf("Exp 10, found %d", bal)
				}
				time.Sleep(1 * time.Millisecond)
			}
		}()
	}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		if bal := bm.GetBalance(zero, qt); bal != 10 {
			b.Errorf("exp 10, found %d", bal)
		}
	}
	b.StopTimer()
	bm.Close()

	for runtime.NumGoroutine() != n {
		time.Sleep(20 * time.Millisecond)
		//fmt.Printf("%d - %d\n", runtime.NumGoroutine(), n)
	}
}

// benchmarkBalanceMapEmptyGets tests against an empty map, focusing on the channel + query speeds
func benchmarkBalanceMapEmptyGets(b *testing.B) {
	n := runtime.NumGoroutine()
	bm := state.NewBalanceMap()
	bm.SetDirectBalance(zero, 10)
	go bm.Serve()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		if bal := bm.GetBalance(zero, qt); bal != 10 {
			b.Errorf("exp 10, found %d", bal)
		}
	}
	b.StopTimer()
	bm.Close()

	for runtime.NumGoroutine() != n {
		//fmt.Printf("%d - %d\n", runtime.NumGoroutine(), n)
		time.Sleep(20 * time.Millisecond)
	}
}

func benchmarkBalanceMapDirect(b *testing.B) {
	n := runtime.NumGoroutine()
	bm := state.NewBalanceMap()
	bm.SetDirectBalance(zero, 10)
	go bm.Serve()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		if bal := bm.GetDirectBalance(zero); bal != 10 {
			b.Errorf("exp 10, found %d", bal)
		}
	}
	b.StopTimer()
	bm.Close()
	for runtime.NumGoroutine() != n {
		time.Sleep(20 * time.Millisecond)
		//fmt.Printf("%d - %d\n", runtime.NumGoroutine(), n)
	}
}

/*
 * Investigation into golang channel performance
 */

var (
	bals      map[int]int64
	totalBals = 1000

	chanPool  chan chan int64
	totalPool = 100
)

func initBals() {
	bals = make(map[int]int64)
	for i := 0; i < totalBals; i++ {
		bals[i] = random.RandInt64Between(0, 50000)
	}
}

func initPool() {
	chanPool = make(chan chan int64, totalPool)
	for i := 0; i < totalPool; i++ {
		chanPool <- make(chan int64, 1)
	}
}

func BenchmarkHandingChannelReturnsSeries(b *testing.B) {
	initBals()
	initPool()
	b.Run("[Series] Direct Map access", benchmarkDirectMap)
	b.Run("[Series] Creating New Channels", benchmarkChannelCreates)
	b.Run("[Series] Pooled Channels", benchmarkChannelPool)
}

func benchmarkDirectMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r := directMap(i % len(bals))
		var _ = r
	}
}

func benchmarkChannelCreates(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r := createChannelForReturn(i % len(bals))
		<-r
	}
}

func benchmarkChannelPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r := useChannelPoolForReturn(i % len(bals))
		var _ = r
	}
}

func createChannelForReturn(addr int) <-chan int64 {
	c := make(chan int64, 1)
	c <- bals[addr]
	//close(c)
	return c
}

func useChannelPoolForReturn(addr int) int64 {
	// Grab the pool
	c := <-chanPool // Caller
	c <- bals[addr] // Result thread

	// Grab the value from the pool resp
	r := <-c      // Caller
	chanPool <- c // Caller
	return r
}

func directMap(addr int) int64 {
	return bals[addr]
}
