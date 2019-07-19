package benchmarks_test

import (
	"testing"

	"github.com/FactomProject/factomd/common/interfaces"

	"github.com/FactomProject/factomd/common/messages"
)

func BenchmarkTypeSwitch(b *testing.B) {
	o := new(messages.Bounce)
	oI := interfaces.IMsg(o)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		switch oI.(type) {
		case interfaces.IMsg:
		default:
			panic("should never happen")
		}
	}
}
