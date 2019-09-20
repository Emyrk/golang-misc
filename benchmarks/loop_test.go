package benchmarks_test

import "testing"

func BenchmarkLoop(b *testing.B) {
	b.Run("Outside", benchmarkOutsideScopeLoop)
	b.Run("Inside", benchmarkInsideScopeLoop)
}

func benchmarkOutsideScopeLoop(b *testing.B) {
	c := int64(0)

	loop := func() {
		for c = 0; c >= 0; c++ {
			if c%100 == 0 {
				break
			}
		}
	}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		loop()
	}
}

func benchmarkInsideScopeLoop(b *testing.B) {
	loop := func() {
		for c := 0; c >= 0; c++ {
			if c%100 == 0 {
				break
			}
		}
	}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		loop()
	}
}
