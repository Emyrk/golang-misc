package main

import (
	"fmt"
)

func main() {
	b1, b2, b3 := [1]byte{}, make([]byte, 1), []byte{0x00}
	fmt.Printf("b1: %#v, b2: %#v", b1, b2)

	// r1, r2 := [1]rune{}, make([]rune, 1)
	// fmt.Printf("r1: %#v, r2: %#v", r1, r2)
}
