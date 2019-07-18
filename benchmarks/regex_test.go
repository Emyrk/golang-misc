package benchmarks_test

import (
	"fmt"
	"regexp"
	"strings"
	"testing"

	"github.com/Emyrk/golang-misc/util"
)

var haystack string
var haystackBytes []byte
var needle string

var realisticHaystackString = "2049799 14:10:03.160    8195-:-0 FollowerExecute                        M-be9992|R-c349a8|H-be9992|0xc022b38780                        EOM[ 0]:   EOM-  DBh/VMh/h 8195/2/-- minute 0 FF  0 --Leader[455b7b] hash[be9992] local"
var realisticHaystackBytes = []byte(realisticHaystackString)
var realisticNeedle = "455b7b"

func init() {
	haystackBytes = util.RandomByteSliceOfLen(50)
	haystack = fmt.Sprintf("%x", haystackBytes)

	needle = haystack[80:86]
}

func BenchmarkRegex(b *testing.B) {
	r, _ := regexp.Compile("[a-f0-9]6")
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.Find(realisticHaystackBytes)
	}
}

func BenchmarkContains(b *testing.B) {
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		strings.Contains(realisticNeedle, realisticHaystackString)
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
