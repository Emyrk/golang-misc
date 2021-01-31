package benchmarks_test

import (
	"sync"
	"testing"
)

var (
	B = true
)

func BenchmarkMutex(b *testing.B) {
	var s sync.RWMutex
	var _ = s
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		s.RLock()
		x := boo()
		var _ = x
		s.RUnlock()
	}
}


func boo() bool {
	return B
}