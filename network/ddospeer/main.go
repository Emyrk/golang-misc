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

	"github.com/FactomProject/factom"
	"github.com/FactomProject/factomd/common/interfaces"
	"github.com/FactomProject/factomd/common/messages"
	"github.com/FactomProject/factomd/common/primitives"
	"github.com/FactomProject/factomd/common/primitives/random"
	"github.com/FactomProject/factomd/p2p"
)

// DefKey is the default signing key
var DefKey *primitives.PrivateKey

// Default identity
var DefID interfaces.IHash

// File where errors go
var ErrorFile *os.File

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

func main() {
	var (
		//192.168.1.10:8108
		peer      = flag.String("addr", "192.168.1.10:8108", "Address to connect to")
		n         = flag.Int("n", 100, "Number of connections per peer")
		spam      = flag.Bool("spam", false, "Enable the spam of messages")
		persecond = flag.Int("ps", 100, "Spam per second. 1000 is max, keep in mind this isn't completely accurate")
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

	// Key setup
	DefKey, _ = primitives.NewPrivateKeyFromHex("4c38c72fc5cdad68f13b74674d3ffb1f3d63a112710868c9b08946553448d26d")
	DefID, _ = primitives.HexToHash("38bab1455b7bd7e5efd15c53c777c79d0c988e9210f1da49a99d95b3a6417be9")

	for _, peerAddr := range peerStrings {
		peers := make([]*BadPeer, *n)
		for i := 0; i < len(peers); i++ {
			random := make([]byte, 10)
			rand.Read(random)
			randomID = fmt.Sprintf("%x", random)
			peers[i] = NewBadPeer(randomID, peerAddr)
			peers[i].Stats = s
			peers[i].StartBadPeer(*spam, *persecond)
		}
	}

	last := time.Now()
	lastSent := int32(0)
	for {
		time.Sleep(2 * time.Second)
		since := time.Since(last)
		sent := s.ReadSentInt()
		fmt.Printf("Conns: %d, Mess: %d, Dials: %d, Sent: %d, Sent %2f/s\n", s.ReadConnInt(), s.ReadMessInt(), s.ReadDialsInt(), sent, float64(sent-lastSent)/float64(since.Seconds()))
		lastSent = sent
		last = time.Now()
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
func (b *BadPeer) StartBadPeer(spam bool, ps int) {
	b.maintainConnection()
	go b.alwaysRead()
	if spam {
		go b.spam(ps)
	}
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

func (b *BadPeer) spam(ps int) {
	if ps > 1000 {
		ps = 1000
	}
	arr := strings.Split(b.Addr, ":")
	factom.SetFactomdServer(arr[0] + ":8088")
	enc := gob.NewEncoder(b.Connection)
	ticker := time.NewTicker(time.Duration(1000/ps) * time.Millisecond)
	for _ = range ticker.C {
		h, err := factom.GetHeights()
		if err != nil {
			//time.Sleep(100 * time.Millisecond)
			continue
		}
		height := h.LeaderHeight

		par := makeAck(height)
		err = enc.Encode(par)
		if err != nil {
			//time.Sleep(100 * time.Millisecond)
			continue
		} else {
			b.Stats.IncSent()
		}
	}
}

func makeAck(height int64) *p2p.Parcel {
	vmIndex := 0

	ack := new(messages.Ack)
	ack.DBHeight = uint32(height) - 1
	ack.VMIndex = vmIndex
	ack.Minute = byte(random.RandIntBetween(0, 9))
	ack.Timestamp = primitives.NewTimestampNow()
	ack.SaltNumber = 0
	copy(ack.Salt[:8], []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00})
	ack.MessageHash = primitives.NewZeroHash()
	ack.LeaderChainID = DefID
	ack.BalanceHash = primitives.NewZeroHash()

	ack.Height = 0
	ack.SerialHash = ack.MessageHash
	ack.Sign(DefKey)

	data, _ := ack.MarshalBinary()

	par := p2p.NewParcel(network, data)
	return par
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
				pong := p2p.NewParcel(network, []byte("Pong"))
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
