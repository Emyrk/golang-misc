package main

import (
	"fmt"
	"os"
	"time"
)

type runner interface {
	run()
}

func main() {
	// after := time.After(time.Millisecond * 5500)
	ticker := time.NewTicker(time.Second * 2)
	defer ticker.Stop()
	ch := time.Tick(time.Second * 3)
	for {
		select {
		case t := <-ticker.C:
			fmt.Println("From ticker:", t)
		case t := <-ch:
			fmt.Println("From chan:", t)
		case <-after(time.Duration(time.Second * time.Duration(5))):
			fmt.Println("timeout, exit")
			os.Exit(1)
		}
	}
}

func after(d time.Duration) <-chan time.Time {
	fmt.Println("NEW")
	return time.After(d)
}
