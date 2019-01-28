package main

import (
	"encoding/gob"
	"flag"
	"io"
	"net"
	"time"

	"strings"

	"fmt"

	"encoding/json"

	"encoding/binary"

	"github.com/FactomProject/factomd/common/constants"
	"github.com/FactomProject/factomd/common/messages"
	"github.com/FactomProject/factomd/common/primitives"
	"github.com/FactomProject/factomd/p2p"
	log "github.com/sirupsen/logrus"
)

var dialer = net.Dialer{
	Timeout:   10 * time.Second,
	KeepAlive: 10 * time.Second,
}

var p2 *Peer

var network p2p.NetworkID

func main() {
	var (
		//192.168.1.10:8108
		peer      = flag.String("addr", "127.0.0.1:8110", "Address to connect to")
		net       = flag.String("net", "LOCAL", "Network you are on")
		customnet = flag.String("customnet", "", "CustomNetID")
	)

	flag.Parse()

	h := primitives.Sha([]byte(*customnet)).Bytes()[:4]
	network = p2p.NetworkID(binary.BigEndian.Uint32(h))

	switch strings.ToLower(*net) {
	case "local":
		network = p2p.LocalNet
	case "custom":
	case "main":
		network = p2p.MainNet
	}

	var _ = customnet
	var _ = peer

	p := NewPeer(*peer)
	p.maintainConnection()
	go p.alwaysRead()
	go p.sends()

	p2 = NewPeer("81.166.138.207:8110")
	p2.maintainConnection()
	go p2.alwaysRead()
	go p2.sends()

	time.Sleep(2 * time.Second)
	fmt.Println("Sending missing dbstate")
	msg := NewMissingDBState(64221, 64224)
	payload, err := msg.MarshalBinary()
	if err != nil {
		log.Error(err)
		return
	}
	replay := p2p.NewParcel(network, payload)
	replay.Header.Type = p2p.TypeMessage

	p.sendParcel(replay)
	time.Sleep(10 * time.Second)

}

func NewMissingDBState(start, end uint32) *messages.DBStateMissing {
	msg := new(messages.DBStateMissing)

	msg.Peer2Peer = true // Always a peer2peer request.
	msg.Timestamp = primitives.NewTimestampNow()
	msg.DBHeightStart = start
	msg.DBHeightEnd = end
	return msg

}

// Peer is our ability to communicate with a factomd
type Peer struct {
	ID         string
	Connection net.Conn
	Addr       string
	dialed     bool
	enc        *gob.Encoder
	dec        *gob.Decoder
	Incoming   chan *P2PParcel
}

func NewPeer(addr string) *Peer {
	p := new(Peer)
	p.Addr = addr
	p.Incoming = make(chan *P2PParcel, 1000)
	return p
}

func (p *Peer) maintainConnection() {
Top:
	//if b.dialed {
	//	b.Stats.IncDials()
	//} else {
	//	b.dialed = true
	//}
	c, err := dialer.Dial("tcp", p.Addr)
	if err != nil {
		time.Sleep(100 * time.Millisecond)
		log.Error(err)
		goto Top
	}
	p.Connection = c
}

func (b *Peer) alwaysRead() {
	b.dec = gob.NewDecoder(b.Connection)
	b.enc = gob.NewEncoder(b.Connection)
	fmt.Println("Reading...")
	for {
		var m p2p.Parcel
		// Reconnect when we can
		err := b.dec.Decode(&m) //bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Error(err)
			if err == io.EOF || err != nil {
				b.maintainConnection()
				b.dec = gob.NewDecoder(b.Connection)
				b.enc = gob.NewEncoder(b.Connection)
			}
			continue
		} else {
			fmt.Println(".", m.Header.Type)
			// always pong!
			if m.Header.Type == p2p.TypePing {
				pong := p2p.NewParcel(network, []byte("Pong"))
				pong.Header.Type = p2p.TypePong
				b.enc.Encode(pong)
			} else if m.Header.Type == p2p.TypeMessage || m.Header.Type == p2p.TypeMessagePart {
				fmt.Println("<")
				replay := p2p.NewParcel(network, m.Payload)
				replay.Header.Type = m.Header.Type
				b.Incoming <- &P2PParcel{replay, time.Time{}}
				//enc.Encode(replay)
			} else if m.Header.Type == p2p.TypePeerRequest {
				list := []p2p.Peer{}
				json, _ := json.Marshal(list)

				response := p2p.NewParcel(network, json)
				response.Header.Type = p2p.TypePeerResponse
				// Send them out to the network - on the connection that requested it!
				b.sendParcel(response)
			}
			// fmt.Printf("Recieved a %s message of %d length payload from %s network.\n", m.MessageType(), len(m.Payload), m.Header.Network.String())
		}
		//fmt.Print("\nMessage from server: " + m.String())
	}
}

func (e *Peer) catch(parcel *P2PParcel) {
	replay := p2p.NewParcel(network, parcel.Parcel.Payload)
	replay.Header.Type = parcel.Parcel.Header.Type
	msg, err := UnmarshalMessage(replay.Payload)
	if err != nil {
		fmt.Println(string(replay.Payload))
		log.Error(err)
		return
	}
	if msg.Type() == constants.DBSTATE_MSG {
		fmt.Println("YES", msg)
		data, _ := msg.MarshalBinary()
		response := p2p.NewParcel(network, data)
		response.Header.Type = p2p.TypeMessagePart
		p2.sendParcel(response)
	} else {
		fmt.Println(msg)
	}
}

func (e *Peer) sends() {
	for {
		select {
		case m := <-e.Incoming:
			cutoff := time.Now().Add(-5 * time.Second)
			if !m.sent.Before(cutoff) {
				time.Sleep(m.sent.Sub(cutoff))
			}

			fmt.Println("Caught")
			e.catch(m)

			// Here is where you could respond to the msg
			if true {
				// Echo?
				e.sendMyParcel(m)
				m.sent = time.Now()
			}
		}
	}
}

func (p *Peer) sendMyParcel(parcel *P2PParcel) {
	replay := p2p.NewParcel(network, parcel.Parcel.Payload)
	replay.Header.Type = parcel.Parcel.Header.Type
	msg, _ := UnmarshalMessage(replay.Payload)

	err := p.enc.Encode(replay)
	if err != nil {
		log.Error(err)
		return
	}

	if msg != nil {
		mt := constants.MessageName(msg.Type())
		fmt.Println("Sent!", p2p.CommandStrings[replay.Header.Type], mt, msg.GetHash().String())
	}
	return
}

func (p *Peer) sendParcel(parcel *p2p.Parcel) {
	err := p.enc.Encode(parcel)
	if err != nil {
		log.Error(err)
		return
	}
}

type P2PParcel struct {
	Parcel *p2p.Parcel
	sent   time.Time
}

func NewP2Parcel(p p2p.Parcel) *P2PParcel {
	pp := new(P2PParcel)
	pp.Parcel = &p
	return pp
}
