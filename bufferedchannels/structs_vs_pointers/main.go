// Comparing memory usage between a struct vs a pointer in a channel

package main

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"time"
	"unsafe"

	"github.com/pkg/profile"
)

var _ = fmt.Sprint
var _ = time.Now()
var _ = debug.FreeOSMemory

type SmallStruct struct {
	Array [25]byte
}

type MediumStruct struct {
	Array [100]byte
}

type LargeStruct struct {
	Array [500]byte
}

func main() {
	runtime.MemProfileRate = 1
	p := profile.Start(profile.MemProfileRate(1), profile.ProfilePath("."))
	defer p.Stop()

	fmt.Printf("SmallSize: %d\nMediumSize:%d\nLargeSize:%d\n", int(unsafe.Sizeof(SmallStruct{})),
		int(unsafe.Sizeof(MediumStruct{})),
		int(unsafe.Sizeof(LargeStruct{})))

	a := make(chan SmallStruct, 100000)
	c := make(chan SmallStruct, 1)
	b := make(chan *SmallStruct, 100000)

	d := make(chan MediumStruct, 100000)
	e := make(chan *MediumStruct, 100000)
	f := make(chan MediumStruct, 1)

	g := make(chan LargeStruct, 100000)
	h := make(chan *LargeStruct, 100000)
	i := make(chan LargeStruct, 1)

	//debug.FreeOSMemory()
	runtime.GC()

	var _, _, _, _, _, _, _, _, _ = a, b, c, d, e, f, g, h, i
	fmt.Println(a, b, c, d, e, f, g, h, i)
}
