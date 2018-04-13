package benchmarks_test

import (
	"testing"
)

var numVals = 3

func BenchmarkMapAllocate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m := make(map[int]int)
		for i := 0; i < numVals; i++ {
			m[i] = i
		}
	}
}

func BenchmarkMapClear(b *testing.B) {
	m := make(map[int]int)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for i := 0; i < numVals; i++ {
			delete(m, i)
		}
		for i := 0; i < numVals; i++ {
			m[i] = i
		}
	}
}
