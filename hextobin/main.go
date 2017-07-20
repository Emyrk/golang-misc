package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"os"
)

func main() {
	h := os.Args[1]

	b, err := hex.DecodeString(h)
	if err != nil {
		panic(err)
	}

	f, err := os.OpenFile("bin.raw", os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		panic(err)
	}
	f.Write(b)

	s := sha512.Sum512(b)
	b = append(s[:], b[:]...)

	s2 := sha256.Sum256(b)
	fmt.Printf("%x\n", s2[:])

}
