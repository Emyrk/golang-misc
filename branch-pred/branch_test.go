package branch_pred_test

import (
	"math/rand"
	"sort"
	"testing"
)

const amt = 100000

func BenchmarkSorted(b *testing.B) {
	array := make([]uint8, b.N)
	for i := range array {
		array[i] = uint8(rand.Int31())
	}

	sort.Slice(array, func(i, j int) bool {
		return array[i] < array[j]
	})

	var sum int64

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if array[i] > uint8(64) {
			sum += int64(array[i])
		}
	}
}

func BenchmarkUnSorted(b *testing.B) {
	array := make([]uint8, b.N)
	for i := range array {
		array[i] = uint8(rand.Int31())
	}

	var sum int64

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if array[i] > uint8(64) {
			sum += int64(array[i])
		}
	}
}
