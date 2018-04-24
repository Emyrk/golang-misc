package benchmarks_test

import (
	"testing"
)

func BenchmarkMapToList(b *testing.B) {
	total := 100
	for i := 0; i < b.N; i++ {
		var m []int
		for i := 0; i < total; i++ {
			m = append(m, i)
		}
	}
}
