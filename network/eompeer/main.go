package main

import (
	"flag"
	"encoding/binary"
	"net"
	"encoding/gob"
	"time"
	"fmt"
	"io"
	"strings"

	"github.com/FactomProject/factomd/p2p"
	"github.com/FactomProject/factomd/common/primitives"
	"github.com/FactomProject/factomd/common/messages"
	"github.com/FactomProject/factomd/common/constants"
)

var network p2p.NetworkID

var dialer = net.Dialer{
	Timeout:   10 * time.Second,
	KeepAlive: 10 * time.Second,
}

func main() {
	var (
		//192.168.1.10:8108
		peer      = flag.String("peers", "localhost:8110", "Address to connect to")
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

	node := NewRepeatingPeer(peerStrings, *n)
	for {
		time.Sleep(1 * time.Second)
	}

	var _ = node

}

type RepeatingPeer struct {
	Connections   []*SingleConnection
	echoWaterfall *EchoWaterfall
}

func NewRepeatingPeer(peers []string, replays int) *RepeatingPeer {
	b := new(RepeatingPeer)

	e := NewEchoWaterFall(nil, b)
	go e.Run()
	for i := 0; i < replays; i++ {
		e = NewEchoWaterFall(e, b)
		go e.Run()
	}
	b.echoWaterfall = e

	for _, p := range peers {
		s := NewConnection(p, b.echoWaterfall.Incoming)
		b.Connections = append(b.Connections, s)
		s.Run()
	}

	return b
}

type EchoWaterfall struct {
	Next     *EchoWaterfall
	Incoming chan *P2PParcel
	B        *RepeatingPeer
}

func NewEchoWaterFall(prev *EchoWaterfall, b *RepeatingPeer) *EchoWaterfall {
	e := new(EchoWaterfall)
	e.Next = prev
	e.B = b
	e.Incoming = make(chan *P2PParcel, 1000)
	return e
}

func (e *EchoWaterfall) Run() {
	for {
		select {
		case m := <-e.Incoming:
			cutoff := time.Now().Add(-5 * time.Second)
			if !m.sent.Before(cutoff) {
				time.Sleep(m.sent.Sub(cutoff))
			}
			e.B.Broadcast(m.Parcel, m.Number)
			m.sent = time.Now()
			m.Number++
			if e.Next != nil {
				e.Next.Incoming <- m
			}
		}
	}
}

func (r *RepeatingPeer) Broadcast(m *p2p.Parcel, number int) {
	for _, c := range r.Connections {
		err := c.Send(m)
		var _ = err
		//if err != nil {
		//	fmt.Println(err)
		//} else {
		//	msg, _ := messages.UnmarshalMessage(m.Payload)
		//	fmt.Println("Msg Sent", number, msg.GetHash().String(), msg.String())
		//}
	}
	msg, _ := messages.UnmarshalMessage(m.Payload)
	fmt.Println("Msg Sent", number, msg.GetHash().String(), msg.String())
}

type SingleConnection struct {
	Connection net.Conn
	Address    string
	Encoder    *gob.Encoder
	Decoder    *gob.Decoder
	UpChannel  chan *P2PParcel
}

func (s *SingleConnection) Run() {
	s.MaintainConnection()
	fmt.Println("Connection Established")
	go s.AlwaysRead()
}

func NewConnection(address string, up chan *P2PParcel) *SingleConnection {
	c := new(SingleConnection)
	c.Address = address
	c.UpChannel = up
	//c.Connection = conn
	//c.Encoder = gob.NewEncoder(conn)
	//c.Decoder = gob.NewDecoder(conn)

	return c
}

func (s *SingleConnection) AlwaysRead() {
	for {
		var m p2p.Parcel
		err := s.Decoder.Decode(&m) //bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			//_, fileError := ErrorFile.WriteString(err.Error() + "\n")
			//if fileError != nil {
			//	panic(fileError)
			//}
			if err == io.EOF || err != nil {
				//b.Stats.DecConns()
				s.MaintainConnection()
				//b.dec = gob.NewDecoder(b.Connection)
				//b.enc = gob.NewEncoder(b.Connection)
			}
			continue
		} else {
			//b.Stats.IncMess()
			if m.Header.Type == p2p.TypePing {
				pong := p2p.NewParcel(network, []byte("Pong"))
				pong.Header.Type = p2p.TypePong
				s.Encoder.Encode(pong)
			} else {

				msg, _ := messages.UnmarshalMessage(m.Payload)
				if msg != nil && msg.Type() == constants.MISSING_MSG {
					continue
				}

				if msg != nil && (msg.Type() == constants.EOM_MSG ||
					msg.Type() == constants.DIRECTORY_BLOCK_SIGNATURE_MSG ||
						true) {
					fmt.Println(len(s.UpChannel), msg.GetHash().String(), msg.String())
					replay := p2p.NewParcel(network, m.Payload)
					replay.Header.Type = p2p.TypeMessage // m.Header.Type
					s.UpChannel <- NewP2Parcel(replay)
				}
				//enc.Encode(replay)
			}
			// fmt.Printf("Recieved a %s message of %d length payload from %s network.\n", m.MessageType(), len(m.Payload), m.Header.Network.String())
		}
		//fmt.Print("\nMessage from server: " + m.String())
	}
}

func (s *SingleConnection) MaintainConnection() {
Top:
	c, err := dialer.Dial("tcp", s.Address)
	if err != nil {
		fmt.Println(err)
		time.Sleep(100 * time.Millisecond)
		goto Top
	}
	s.Connection = c
	s.Encoder = gob.NewEncoder(c)
	s.Decoder = gob.NewDecoder(c)
}

func (s *SingleConnection) Send(m *p2p.Parcel) error {
	if s.Encoder == nil {
		return fmt.Errorf("No encoder")
	}
	return s.Encoder.Encode(m)
}
