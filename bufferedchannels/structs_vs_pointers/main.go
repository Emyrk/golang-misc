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

	sa := make(chan SmallStruct, 100000)
	sb := make(chan *SmallStruct, 100000)
	sc := make(chan SmallStruct, 1)
	sd := make(chan *SmallStruct, 1)

	ma := make(chan MediumStruct, 100000)
	mb := make(chan *MediumStruct, 100000)
	mc := make(chan MediumStruct, 1)
	md := make(chan *MediumStruct, 1)

	la := make(chan LargeStruct, 100000)
	lb := make(chan *LargeStruct, 100000)
	lc := make(chan LargeStruct, 1)
	ld := make(chan *LargeStruct, 1)

	//debug.FreeOSMemory()
	runtime.GC()

	fmt.Println(sa, sb, sc, sd,
		ma, mb, mc, md,
		la, lb, lc, ld,
	)
}
