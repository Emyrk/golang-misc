package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {

	// Given a set of difficulties
	difficulties := randDiffs(1000)
	sort.SliceStable(difficulties, func(i, j int) bool { return difficulties[i] > difficulties[j] })
	percentDiffs(difficulties)

}

func percentDiffs(arr []uint64) {
	for i, v := range arr {
		if i == 0 {
			continue
		}

		p := float64(v) / float64(arr[i-1])

		fmt.Printf("%d %.3f\n", i, p)
	}
}

func hashedDiffs() []uint64 {

}

func randDiffs(l int) []uint64 {
	arr := make([]uint64, l)
	for i := range arr {
		arr[i] = rand.Uint64()
	}

	return arr
}
