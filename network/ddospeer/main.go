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

	ed "github.com/FactomProject/ed25519"
	"github.com/FactomProject/factom"
	"github.com/FactomProject/factomd/common/factoid"
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
var Unique bool

func main() {
	var (
		//192.168.1.10:8108
		peer      = flag.String("addr", "192.168.1.10:8108", "Address to connect to")
		n         = flag.Int("n", 100, "Number of connections per peer")
		ackSpam   = flag.Bool("aspam", false, "Enable the spam of messages")
		persecond = flag.Int("ps", 100, "Spam per second. 1000 is max, keep in mind this isn't completely accurate")
		net       = flag.String("net", "LOCAL", "Network you are on")
		customnet = flag.String("customnet", "", "CustomNetID")
		unique    = flag.Bool("u", false, "Unique messages")
	)

	flag.Parse()
	Unique = *unique
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
	DefKey, err = primitives.NewPrivateKeyFromHex("4c38c72fc5cdad68f13b74674d3ffb1f3d63a112710868c9b08946553448d26d")
	DefID, _ = primitives.HexToHash("38bab1455b7bd7e5efd15c53c777c79d0c988e9210f1da49a99d95b3a6417be9")

	if err != nil {
		panic(err)
	}
	for _, peerAddr := range peerStrings {
		peers := make([]*BadPeer, *n)
		for i := 0; i < len(peers); i++ {
			random := make([]byte, 10)
			rand.Read(random)
			randomID = fmt.Sprintf("%x", random)
			peers[i] = NewBadPeer(randomID, peerAddr)
			peers[i].Stats = s
			peers[i].StartBadPeer(*ackSpam, *persecond)
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
func (b *BadPeer) StartBadPeer(ackSpam bool, ps int) {
	b.maintainConnection()
	go b.alwaysRead()
	if ackSpam {
		go b.ackSpam(ps)
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

func (b *BadPeer) ackSpam(ps int) {
	if ps > 1000 {
		ps = 1000
	}
	ps = ps * 1000
	tickertime := time.Duration(1e12/ps) * time.Nanosecond
	arr := strings.Split(b.Addr, ":")
	factom.SetFactomdServer(arr[0] + ":8088")
	enc := gob.NewEncoder(b.Connection)
	ticker := time.NewTicker(tickertime)

	a := new(AckMaker)
	last := time.Now().Add(-20 * time.Second)
	var height int64
	for _ = range ticker.C {
		if time.Since(last).Seconds() > 10 {
			h, err := factom.GetHeights()
			if err != nil {
				//time.Sleep(100 * time.Millisecond)
				continue
			}
			height = h.LeaderHeight
			last = time.Now()
		}

		par := a.makeAck(height)
		//msgs := a.makeValidAckWithMessage(height)
		err := enc.Encode(par)
		if err != nil {
			//time.Sleep(100 * time.Millisecond)
			continue
		} else {
			b.Stats.IncSent()
		}

		//err = enc.Encode(msgs[1])
		//if err != nil {
		//	//time.Sleep(100 * time.Millisecond)
		//	continue
		//} else {
		//	b.Stats.IncSent()
		//}
	}
}

func (a *AckMaker) makeValidAckWithMessage(height int64) []*p2p.Parcel {

	if a.Ack == nil || a.Ack.Height < uint32(height) || Unique {
		msg, err := fctTrans(100)
		if err != nil {
			panic(err)
		}

		data, _ := msg.MarshalBinary()
		par := p2p.NewParcel(network, data)
		a.M = par

		a.makeAckForMsg(height, msg)
	}

	return []*p2p.Parcel{a.P, a.M}
}

// AckMaker makes acks
type AckMaker struct {
	Ack     *messages.Ack
	AckData []byte

	P *p2p.Parcel
	M *p2p.Parcel
}

func (a *AckMaker) makeAck(height int64) *p2p.Parcel {
	if a.Ack == nil || a.Ack.Height < uint32(height) || Unique {
		msg := new(messages.Bounce)
		msg.Timestamp = primitives.NewTimestampNow()
		a.makeAckForMsg(height, msg)
	}
	return a.P
}

func (a *AckMaker) makeAckForMsg(height int64, msg interfaces.IMsg) *p2p.Parcel {
	if a.Ack == nil || a.Ack.Height < uint32(height) || Unique {
		vmIndex := 0

		h := height - 1
		if h < 0 {
			h = 0
		}
		ack := new(messages.Ack)
		ack.DBHeight = uint32(h)
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

		a.Ack = ack
		a.Ack.Height = 200
		a.Ack.DBHeight = uint32(height)
		ack.Sign(DefKey)
		data, _ := ack.MarshalBinary()
		a.AckData = data
		par := p2p.NewParcel(network, a.AckData)
		a.P = par
	}

	return a.P
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

func fctTrans(amt uint64) (interfaces.IMsg, error) {
	inSec, _ := primitives.HexToHash("FB3B471B1DCDADFEB856BD0B02D8BF49ACE0EDD372A3D9F2A95B78EC12A324D6")
	outEC, _ := primitives.HexToHash("c23ae8eec2beb181a0da926bd2344e988149fbe839fbc7489f2096e7d6110243")
	inHash, _ := primitives.HexToHash("646F3E8750C550E4582ECA5047546FFEF89C13A175985E320232BACAC81CC428")
	var sec [64]byte
	copy(sec[:32], inSec.Bytes())

	pub := ed.GetPublicKey(&sec)
	//inRcd := shad(inPub.Bytes())

	rcd := factoid.NewRCD_1(pub[:])
	inAdd := factoid.NewAddress(inHash.Bytes())
	outAdd := factoid.NewAddress(outEC.Bytes())

	trans := new(factoid.Transaction)
	trans.AddInput(inAdd, amt)
	trans.AddECOutput(outAdd, amt)

	trans.AddRCD(rcd)
	trans.AddAuthorization(rcd)
	trans.SetTimestamp(primitives.NewTimestampNow())

	fee, err := trans.CalculateFee(4000)
	if err != nil {
		return nil, err
	}
	input, err := trans.GetInput(0)
	if err != nil {
		return nil, err
	}
	input.SetAmount(amt + fee)

	dataSig, err := trans.MarshalBinarySig()
	if err != nil {
		return nil, err
	}
	sig := factoid.NewSingleSignatureBlock(inSec.Bytes(), dataSig)
	trans.SetSignatureBlock(0, sig)

	msg := new(messages.FactoidTransaction)
	msg.Transaction = trans

	return msg, nil
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
