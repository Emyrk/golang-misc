package main

import (
	"crypto/rand"
	"encoding/gob"
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
	"sync/atomic"
	"time"

	"github.com/FactomProject/factomd/p2p"
)

var ErrorFile *os.File

type Message struct {
	Message string
}

var dialer = net.Dialer{
	Timeout:   10 * time.Second,
	KeepAlive: 10 * time.Second,
}

var conn net.Conn
var randomID string

func main() {
	var (
		//192.168.1.10:8108
		peer = flag.String("addr", "192.168.1.10:8108", "Address to connect to")
	)

	flag.Parse()
	peerStrings := strings.Split(*peer, " ")

	s := NewStatMaintainer()
	os.Remove("errs.txt")
	errFile, err := os.OpenFile("errs.txt", os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		panic(err)
	}
	ErrorFile = errFile
	ErrorFile.WriteString("Error file\n")

	for _, peerAddr := range peerStrings {
		peers := make([]*BadPeer, 100)
		for i := 0; i < len(peers); i++ {
			random := make([]byte, 10)
			rand.Read(random)
			randomID = fmt.Sprintf("%x", random)
			peers[i] = NewBadPeer(randomID, peerAddr)
			peers[i].Stats = s
			peers[i].StartBadPeer()
		}
	}

	for {
		time.Sleep(2 * time.Second)
		fmt.Printf("Conns: %d, Mess: %d, Dials: %d\n", s.ReadConnInt(), s.ReadMessInt(), s.ReadDialsInt())
	}
}

// BadPeer keeps dialing
type BadPeer struct {
	ID         string
	Connection net.Conn
	Addr       string
	Stats      *StatMaintainer
	dialed     bool
}

// NewBadPeer inits
func NewBadPeer(id string, addr string) *BadPeer {
	b := new(BadPeer)
	b.ID = id
	b.Addr = addr

	return b
}

// StartBadPeer runs
func (b *BadPeer) StartBadPeer() {
	b.maintainConnection()
	go b.alwaysRead()
}

func (b *BadPeer) maintainConnection() {
Top:
	if b.dialed {
		b.Stats.IncDials()
	} else {
		b.dialed = true
	}
	c, err := dialer.Dial("tcp", b.Addr)
	if err != nil {
		time.Sleep(100 * time.Millisecond)
		goto Top
	}
	b.Connection = c
	b.Stats.IncConns()
}

func (b *BadPeer) alwaysRead() {
	dec := gob.NewDecoder(b.Connection)
	for {
		var m p2p.Parcel
		err := dec.Decode(&m) //bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			_, fileError := ErrorFile.WriteString(err.Error() + "\n")
			if fileError != nil {
				panic(fileError)
			}
			if err == io.EOF || err != nil {
				b.Stats.DecConns()
				b.maintainConnection()
				dec = gob.NewDecoder(b.Connection)
			}
			continue
		} else {
			b.Stats.IncMess()
			if m.Header.Type == p2p.TypePing {
				pong := p2p.NewParcel(p2p.LocalNet, []byte("Pong"))
				pong.Header.Type = p2p.TypePong
				enc := gob.NewEncoder(b.Connection)
				enc.Encode(pong)
			}
			// fmt.Printf("Recieved a %s message of %d length payload from %s network.\n", m.MessageType(), len(m.Payload), m.Header.Network.String())
		}
		//fmt.Print("\nMessage from server: " + m.String())
	}
}

// StatMaintainer maintains some stat
type StatMaintainer struct {
	conns *int32
	mess  *int32
	dials *int32
}

// NewStatMaintainer init
func NewStatMaintainer() *StatMaintainer {
	s := new(StatMaintainer)
	s.conns = new(int32)
	s.mess = new(int32)
	s.dials = new(int32)
	return s
}

func (s *StatMaintainer) ReadConnInt() int32 {
	k := atomic.LoadInt32(s.conns)
	return k
}

// IncConns ++
func (s *StatMaintainer) IncConns() {
	atomic.AddInt32(s.conns, 1)
}

// DecConns --
func (s *StatMaintainer) DecConns() {
	atomic.AddInt32(s.conns, -1)
}

func (s *StatMaintainer) ReadMessInt() int32 {
	k := atomic.LoadInt32(s.mess)
	return k
}

// IncConns ++
func (s *StatMaintainer) IncMess() {
	atomic.AddInt32(s.mess, 1)
}

// DecConns --
func (s *StatMaintainer) DecMess() {
	atomic.AddInt32(s.mess, -1)
}

func (s *StatMaintainer) ReadDialsInt() int32 {
	k := atomic.LoadInt32(s.dials)
	return k
}

// IncConns ++
func (s *StatMaintainer) IncDials() {
	atomic.AddInt32(s.dials, 1)
}

// DecConns --
func (s *StatMaintainer) DecDials() {
	atomic.AddInt32(s.dials, -1)
}
