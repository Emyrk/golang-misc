package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// RandomStruct is just a struct
type RandomStruct struct {
	Value int
}

var arrayPointer []*RandomStruct
var end *int32

var group sync.WaitGroup

func main() {
	end = new(int32)
	arrayPointer = make([]*RandomStruct, 100000)
	arrayPointer[0] = &RandomStruct{Value: 1}

	// Return when both routines done
	group.Add(2)
	go appendArray()
	go readArray()
	group.Wait()
	fmt.Println("Done!")
}

func appendArray() {
	// write all elements
	for i := 0; i < 100000; i++ {
		time.Sleep(1 * time.Nanosecond)
		arrayPointer[i] = &RandomStruct{Value: i}
		//arrayPointer = append(arrayPointer, &RandomStruct{Value: i})
		atomic.AddInt32(end, 1)
	}
	group.Done()
	fmt.Println("Write Done")
}

func readArray() {
	for i := 0; i < 100000; i++ {
		// Wait until the next element is written
		for i >= int(atomic.LoadInt32(end)) {
			// Allow atomic to write with small sleep
			time.Sleep(1 * time.Nanosecond)
		}
		if arrayPointer[i].Value != i {
			panic(fmt.Sprintf("Value wrong. Exp %d, found %d", i, arrayPointer[i].Value))
		}
		if i%1000 == 0 {
			fmt.Println(arrayPointer[i].Value)
		}
	}
	group.Done()
	fmt.Println("Read Done")
}
