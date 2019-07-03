package benchmarks_test

import (
	"testing"
)

func emptyOp() bool {
	return true
}

func BenchmarkBoolCompare(b *testing.B) {
	comp := false
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		if comp {
			emptyOp()
		} else {
			emptyOp()
		}
	}
}

func BenchmarkEmptyStringCompare(b *testing.B) {
	comp := ""
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		if comp == "" {
			emptyOp()
		}
	}
}

//
//func TestRegex(t *testing.T) {
//	r, _ := regexp.Compile("test")
//	s := []byte("BLAHtest")
//	var _, _ = s, r
//}

// func BenchmarkFoo(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		fmt.Println("TEST")
// 		// perform the operation we're analyzing
// 	}
// }

// 100 char
// last half
// 6 char hex
