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
var block sync.RWMutex
var end *int32

var group sync.WaitGroup

func main() {
	end = new(int32)
	arrayPointer = make([]*RandomStruct, 100000)
	arrayPointer[0] = &RandomStruct{Value: 1}
	block.Lock()
	group.Add(2)
	go appendArray()
	go readArray()
	block.Unlock()
	time.Sleep(1 * time.Second)
	group.Wait()
	fmt.Println("Done!")
}

func appendArray() {
	block.RLock()
	for i := 0; i < 100000; i++ {
		time.Sleep(1 * time.Nanosecond)
		arrayPointer[i] = &RandomStruct{Value: i}
		//arrayPointer = append(arrayPointer, &RandomStruct{Value: i})
		//end = i
		atomic.AddInt32(end, 1)
	}
	block.RUnlock()
	group.Done()
	fmt.Println("Write Done")
}

func readArray() {
	block.RLock()
	for i := 0; i < 100000; i++ {
		for i >= int(atomic.LoadInt32(end)) {

		}
		if arrayPointer[i].Value != i {
			panic(fmt.Sprintf("Value wrong. Exp %d, found %d", i, arrayPointer[i].Value))
		}
		fmt.Println(arrayPointer[i].Value)
	}
	block.RUnlock()
	group.Done()
	fmt.Println("Read Done")
}
