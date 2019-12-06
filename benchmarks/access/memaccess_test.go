package benchmarks_test

import (
	crand "crypto/rand"
	"math/rand"
	"testing"
)

var RandomSlice []byte

func BenchmarkMemoryAccess(b *testing.B) {
	RandomSlice = make([]byte, 1024*1024*100)
	crand.Read(RandomSlice)

	b.Run("Sequential", benchmarkSequentialAccess)
	b.Run("Random", benchmarkRandomAccess)
	b.Run("Random", benchmarkRandomAccessCast)
}

func benchmarkSequentialAccess(b *testing.B) {
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		index := rand.Int()
		// noop the index
		v := RandomSlice[i%len(RandomSlice)]
		index += i + int(v)
	}
}

func benchmarkRandomAccess(b *testing.B) {
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		index := rand.Int()
		// noop the index
		var _ = RandomSlice[index%len(RandomSlice)]
		//index += i + int(v)
	}
}

func benchmarkRandomAccessCast(b *testing.B) {
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		index := rand.Int()
		// noop the index
		v := RandomSlice[index%len(RandomSlice)]
		index += i + int(v)
	}
}
