package main

import (
	crand "crypto/rand"
	"crypto/sha256"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"sync"
	"time"

	"github.com/pegnet/LXRHash"
)

func main() {
	a := make(chan int, 100)
	populate(a)
	close(a)

	runtime.GOMAXPROCS(1)
	for i := 0; i < 5; i++ {
		go lxSpin()
	}

	runtime.Gosched()
	// Time how long it takes to drain in two scenarios
	drainFor(a)
}

func populate(c chan int) {
	for len(c) != cap(c) {
		c <- len(c)
	}
}

func spin() {
	var i int
	for {
		i++
	}
}

func hashSpin() {
	r := make([]byte, 32)
	_, _ = crand.Read(r)
	d := sha256.Sum256(r)

	for {
		d = sha256.Sum256(d[:])
	}
}

// LX holds an instance of lxrhash
var LX lxr.LXRHash
var lxInitializer sync.Once

// The init function for LX is expensive. So we should explicitly call the init if we intend
// to use it. Make the init call idempotent
func InitLX() {
	lxInitializer.Do(func() {
		// This code will only be executed ONCE, no matter how often you call it
		LX.Verbose(true)
		if size, err := strconv.Atoi(os.Getenv("LXRBITSIZE")); err == nil && size >= 8 && size <= 30 {
			LX.Init(0xfafaececfafaecec, uint64(size), 256, 5)
		} else {
			LX.Init(lxr.Seed, lxr.MapSizeBits, lxr.HashSize, lxr.Passes)
		}
	})
}

func lxSpin() {
	InitLX()
	d := make([]byte, 32)
	_, _ = crand.Read(d)

	for {
		d = LX.Hash(d[:])
	}
}

func drainRange(c chan int) {
	start := time.Now()
	for _ = range c {

	}
	fmt.Printf("Took %s\n", time.Since(start))
}

func drainFor(c chan int) {
	start := time.Now()

MainLoop:
	for {
		select {
		case _, open := <-c:
			if !open {
				break MainLoop
			}
		}
	}
	fmt.Printf("Took %s\n", time.Since(start))
}
