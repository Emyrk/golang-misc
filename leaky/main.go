package main

import (
	"context"
	"fmt"
	"runtime"
)

func main() {

	g := runtime.NumGoroutine()
	fmt.Printf("Starting with %d routines\n", g)

	ctx, cancel := context.WithCancel(context.Background())
	fmt.Printf("1 cancel ctx, # routines = %d\n", g)

	cancel()
	fmt.Printf("After cancel ctx, # routines = %d\n", g)
	<-ctx.Done()
}
