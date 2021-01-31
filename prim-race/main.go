package main

import (
	"fmt"
	"math/rand"
	"time"
)

var prim bool

func main() {
	for i := 0; i < 5; i++ {
		go primChanger(time.Second)
	}

	for {
		time.Sleep(time.Second)
		fmt.Println(prim)
	}
}

func primChanger(d time.Duration) {
	for {
		time.Sleep(time.Duration(rand.Int63n(int64(d))))
		prim = !prim
	}
}
