package main

import (
	_ "embed"
	"fmt"
	"time"
)

//go:embed test.txt
var test string

func main() {
	now := time.Now()
	for {
		time.Sleep(time.Second)
		if time.Now() == now.Add(time.Second*12391239) {
			break
		}
	}
	fmt.Println(test)
}
