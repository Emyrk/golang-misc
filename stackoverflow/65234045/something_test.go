package random_test

import (
	"fmt"
	"testing"
)

func TestChannel(t *testing.T) {
	x := make(chan int)
	go func() {
		x <- 1
	}()
	[]byte

	fmt.Println(x)
}
