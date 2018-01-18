package main

import (
	"crypto/rand"
	"encoding/binary"
	"encoding/gob"
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
	"sync/atomic"
	"time"

	//"github.com/FactomProject/factom"
	//"github.com/FactomProject/factomd/common/interfaces"
	"github.com/FactomProject/factomd/common/messages"
	"github.com/FactomProject/factomd/common/constants"
	//"github.com/FactomProject/factomd/common/primitives"
	//"github.com/FactomProject/factomd/common/primitives/random"
	"github.com/FactomProject/factomd/p2p"
	"github.com/FactomProject/factomd/common/primitives"
)

// Message is something
type Message struct {
	Message string
}

var dialer = net.Dialer{
	Timeout:   10 * time.Second,
	KeepAlive: 10 * time.Second,
}

var conn net.Conn
var randomID string
var network p2p.NetworkID
var Unique bool

// File where errors go
var ErrorFile *os.File

func main() {
	var (
		//192.168.1.10:8108
		peer      = flag.String("addr", "localhost:8110", "Address to connect to")
		n         = flag.Int("n", 100, "Number of replays")
		net       = flag.String("net", "LOCAL", "Network you are on")
		customnet = flag.String("customnet", "", "CustomNetID")
	)

	flag.Parse()
	peerStrings := strings.Split(*peer, " ")
	custnetID := primitives.Sha([]byte(*customnet)).Bytes()[:4]

	network = p2p.LocalNet
	if *net == "CUSTOM" {
		network = p2p.NetworkID(binary.BigEndian.Uint32(custnetID))
	}

	s := NewStatMaintainer()
	os.Remove("errs.txt")
	errFile, err := os.OpenFile("errs.txt", os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		panic(err)
	}
	ErrorFile = errFile
	ErrorFile.WriteString("Error file\n")

	peers := make([]*BadPeer, 1)
	for _, peerAddr := range peerStrings {
		for i := 0; i < len(peers); i++ {
			random := make([]byte, 10)
			rand.Read(random)
			randomID = fmt.Sprintf("%x", random)
			peers[i] = NewBadPeer(randomID, peerAddr, *n)
			peers[i].Stats = s
			peers[i].StartBadPeer(*n)
		}
	}

	//last := time.Now()
	//lastSent := int32(0)
	for {
		time.Sleep(2 * time.Second)
		//since := time.Since(last)
		sent := s.ReadSentInt()
		fmt.Printf("Conns: %d, Mess: %d, Dials: %d, Sent: %d, WaterFall %d\n", s.ReadConnInt(), s.ReadMessInt(), s.ReadDialsInt(), sent, len(peers[0].echoWaterfall.Incoming))
		//lastSent = sent
		//last = time.Now()
	}
}

// BadPeer keeps dialing
type BadPeer struct {
	ID            string
	Connection    net.Conn
	Addr          string
	Stats         *StatMaintainer
	dialed        bool
	echoWaterfall *EchoWaterfall
	enc           *gob.Encoder
	dec           *gob.Decoder
}

// NewBadPeer inits
func NewBadPeer(id string, addr string, replays int) *BadPeer {
	b := new(BadPeer)
	b.ID = id
	b.Addr = addr

	e := NewEchoWaterFall(nil, b)
	go e.Run()
	for i := 0; i < replays; i++ {
		e = NewEchoWaterFall(e, b)
		go e.Run()
	}
	b.echoWaterfall = e

	return b
}

func NewEchoWaterFall(prev *EchoWaterfall, b *BadPeer) *EchoWaterfall {
	e := new(EchoWaterfall)
	e.Next = prev
	e.B = b
	e.Incoming = make(chan *P2PParcel, 1000)
	return e
}

type EchoWaterfall struct {
	Next     *EchoWaterfall
	Incoming chan *P2PParcel
	B        *BadPeer
}

func (e *EchoWaterfall) Run() {
	for {
		select {
		case m := <-e.Incoming:
			cutoff := time.Now().Add(-5 * time.Second)
			if !m.sent.Before(cutoff) {
				time.Sleep(m.sent.Sub(cutoff))
			}
			e.B.echo(m)
			m.sent = time.Now()
			if e.Next != nil {
				e.Next.Incoming <- m
			}
		}
	}
}

type P2PParcel struct {
	Parcel *p2p.Parcel
	sent   time.Time
}

// StartBadPeer runs
func (b *BadPeer) StartBadPeer(replays int) {
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

func (b *BadPeer) echo(parcel *P2PParcel) {
	for {
		replay := p2p.NewParcel(network, parcel.Parcel.Payload)
		replay.Header.Type = parcel.Parcel.Header.Type
		msg, _ := messages.UnmarshalMessage(replay.Payload)

		err := b.enc.Encode(replay)
		if err != nil {
			fmt.Println(err)
			time.Sleep(1 * time.Second)
			continue
		}

		if msg != nil {
			mt := fmt.Sprintf("%d", msg.Type())
			if msg.Type() == constants.EOM_MSG {
				mt = "EOM"
			}
			fmt.Println("Echo!", p2p.CommandStrings[replay.Header.Type], mt, msg.GetHash().String())
		}
		return
	}
}

func (b *BadPeer) alwaysRead() {
	b.dec = gob.NewDecoder(b.Connection)
	b.enc = gob.NewEncoder(b.Connection)
	for {
		var m p2p.Parcel
		err := b.dec.Decode(&m) //bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			_, fileError := ErrorFile.WriteString(err.Error() + "\n")
			if fileError != nil {
				panic(fileError)
			}
			if err == io.EOF || err != nil {
				b.Stats.DecConns()
				b.maintainConnection()
				b.dec = gob.NewDecoder(b.Connection)
				b.enc = gob.NewEncoder(b.Connection)
			}
			continue
		} else {
			b.Stats.IncMess()
			if m.Header.Type == p2p.TypePing {
				pong := p2p.NewParcel(network, []byte("Pong"))
				pong.Header.Type = p2p.TypePong
				b.enc.Encode(pong)
			} else {
				replay := p2p.NewParcel(network, m.Payload)
				replay.Header.Type = m.Header.Type
				b.echoWaterfall.Incoming <- &P2PParcel{replay, time.Time{}}
				//enc.Encode(replay)
			}
			// fmt.Printf("Recieved a %s message of %d length payload from %s network.\n", m.MessageType(), len(m.Payload), m.Header.Network.String())
		}
		//fmt.Print("\nMessage from server: " + m.String())
	}
}

/*
 * Stats
 */
// StatMaintainer maintains some stat
type StatMaintainer struct {
	conns *int32
	mess  *int32
	dials *int32
	sent  *int32
}

// NewStatMaintainer init
func NewStatMaintainer() *StatMaintainer {
	s := new(StatMaintainer)
	s.conns = new(int32)
	s.mess = new(int32)
	s.dials = new(int32)
	s.sent = new(int32)
	return s
}

// ReadConnInt ==
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

// ReadMessInt ==
func (s *StatMaintainer) ReadMessInt() int32 {
	k := atomic.LoadInt32(s.mess)
	return k
}

// IncMess ++
func (s *StatMaintainer) IncMess() {
	atomic.AddInt32(s.mess, 1)
}

// DecMess --
func (s *StatMaintainer) DecMess() {
	atomic.AddInt32(s.mess, -1)
}

// ReadDialsInt ..
func (s *StatMaintainer) ReadDialsInt() int32 {
	k := atomic.LoadInt32(s.dials)
	return k
}

// IncDials ++
func (s *StatMaintainer) IncDials() {
	atomic.AddInt32(s.dials, 1)
}

// DecDials --
func (s *StatMaintainer) DecDials() {
	atomic.AddInt32(s.dials, -1)
}

// ReadSentInt ..
func (s *StatMaintainer) ReadSentInt() int32 {
	k := atomic.LoadInt32(s.sent)
	return k
}

// IncSent ++
func (s *StatMaintainer) IncSent() {
	atomic.AddInt32(s.sent, 1)
}

// DecSent --
func (s *StatMaintainer) DecSent() {
	atomic.AddInt32(s.sent, -1)
}
