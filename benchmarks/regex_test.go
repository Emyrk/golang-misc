package benchmarks_test

import (
	"testing"
	"regexp"
	"strings"
)

func BenchmarkRegex(b *testing.B) {
	r, _ := regexp.Compile(".*test.*")
	s := []byte("BLAHtest")
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.Find(s)
	}
}

func BenchmarkContains(b *testing.B) {
	s := "BLAHtest"
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		strings.Contains(s, "test")
	}
}

func TestRegex(t *testing.T) {
	r, _ := regexp.Compile(".*test.*")
	s := []byte("BLAHtest")
	var _, _ = s, r
}

// func BenchmarkFoo(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		fmt.Println("TEST")
// 		// perform the operation we're analyzing
// 	}
// }
