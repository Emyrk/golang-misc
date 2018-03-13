package benchmarks_test

import (
	"testing"
)

var channel = make(chan int, 1)

func BenchmarkGoRoutine(b *testing.B) {
	for i := 0; i < b.N; i++ {
		go func() {
			return
		}()
	}
}

func BenchmarkChannel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		channel <- i
		<- channel
	}
}
