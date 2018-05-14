package main

import (
	"flag"
	"log"

	"os"

	"fmt"

	"bufio"

	"io/ioutil"
	"net"
	"strconv"

	"sort"
	"strings"

	"bytes"
	"encoding/gob"

	"time"

	"github.com/Emyrk/golang-misc/meshtesting/meshconn"
	"github.com/weaveworks/mesh"
)

func main() {
	peers := &stringset{}
	var (
		//httpListen = flag.String("http", ":8080", "HTTP listen address")
		meshListen = flag.String("mesh", net.JoinHostPort("0.0.0.0", strconv.Itoa(mesh.Port)), "mesh listen address")
		hwaddr     = flag.String("hwaddr", MustHardwareAddr(), "MAC address, i.e. mesh peer ID")
		nickname   = flag.String("nickname", MustHostname(), "peer nickname")
		password   = flag.String("password", "", "password (optional)")
		channel    = flag.String("channel", "default", "gossip channel name")
	)
	flag.Var(peers, "peer", "initial peer (may be repeated)")
	flag.Parse()

	logger := log.New(os.Stderr, *nickname+"> ", log.LstdFlags)

	host, portStr, err := net.SplitHostPort(*meshListen)
	if err != nil {
		logger.Fatalf("mesh address: %s: %v", *meshListen, err)
	}
	port, err := strconv.Atoi(portStr)
	if err != nil {
		logger.Fatalf("mesh address: %s: %v", *meshListen, err)
	}

	name, err := mesh.PeerNameFromString(*hwaddr)
	if err != nil {
		panic(err)
	}
	fmt.Println(port)

	router, err := mesh.NewRouter(mesh.Config{
		Host:               host,
		Port:               port,
		ProtocolMinVersion: mesh.ProtocolMinVersion,
		Password:           []byte(*password),
		ConnLimit:          64,
		PeerDiscovery:      true,
		TrustedSubnets:     []*net.IPNet{},
	}, name, *nickname, mesh.NullOverlay{}, log.New(ioutil.Discard, "", 0))
	func() {
		logger.Printf("mesh router starting (%s)", *meshListen)
		router.Start()
	}()
	defer func() {
		logger.Printf("mesh router stopping")
		router.Stop()
	}()
	router.ConnectionMaker.InitiateConnections(peers.slice(), true)

	fmt.Println(peers.slice())
	peer := meshconn.NewPeer(name, 0, logger)

	gossip, err := router.NewGossip(*channel, peer)
	if err != nil {
		logger.Fatalf("Could not create gossip: %v", err)
	}

	peer.Register(gossip)

	peerlist := func() {
		for _, ap := range router.Peers.Descriptions() {
			fmt.Printf("%s : %d : %s : %d\n", ap.Name, ap.UID, ap.NickName, ap.NumConnections)
		}
	}

	to, _ := mesh.PeerNameFromString("1b:b3:c4:08:33:f2")
	to, _ = mesh.PeerNameFromString("00:00:00:00:00:01")
	//fmt.Println("SEND")
	//_, err = peer.WriteTo([]byte("Hello"), meshconn.MeshAddr{PeerName: to})
	//fmt.Println(err)
	var _ = to

	go Recieve(peer)
	for {
		// read in input from stdin
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")
		text, _ := reader.ReadString('\n')
		// send to socket
		if len(text) == 0 {
			text = "Nothing"
			time.Sleep(1 * time.Second)
		}
		peer.Write([]byte(text))
		fmt.Printf("sent to %d peers\n", len(router.Peers.Descriptions()))
		//peer.OnGossipBroadcast(name, []byte(text))
		//fmt.Fprintf(conn, text+"\n")
		peerlist()
	}

}

/*

type pkt struct {
	SrcName mesh.PeerName
	SrcUID  mesh.PeerUID
	Buf     []byte
}

*/

type Packet struct {
	Data []byte
	To   meshconn.MeshAddr
}

func NewPacket(data []byte) mesh.GossipData {
	p := new(Packet)
	p.Data = data
	return p
}

// Encode serializes our complete state to a slice of byte-slices.
// In this simple example, we use a single gob-encoded
// buffer: see https://golang.org/pkg/encoding/gob/
func (st *Packet) Encode() [][]byte {
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(st.Data); err != nil {
		panic(err)
	}
	return [][]byte{buf.Bytes()}
}

// Merge merges the other GossipData into this one,
// and returns our resulting, complete state.
func (st *Packet) Merge(other mesh.GossipData) (complete mesh.GossipData) {
	return other
}

func Recieve(p *meshconn.Peer) {
	//var n mesh.PeerName
	for {
		b, from, broadcast, err := p.ReadFrom()
		fmt.Println()
		//m, err := p.OnGossip(b)
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Printf("\n[%d] %s: %s", len(b), from.String(), string(string(b)))
		if broadcast {
			p.WriteTo([]byte(fmt.Sprintf("Echo from %s : %s. UID: %d", from, string(b), from.(meshconn.MeshAddr).PeerUID)), from)
		}
	}
}

type stringset map[string]struct{}

func (ss stringset) Set(value string) error {
	ss[value] = struct{}{}
	return nil
}

func (ss stringset) String() string {
	return strings.Join(ss.slice(), ",")
}

func (ss stringset) slice() []string {
	slice := make([]string, 0, len(ss))
	for k := range ss {
		slice = append(slice, k)
	}
	sort.Strings(slice)
	return slice
}

func MustHardwareAddr() string {
	ifaces, err := net.Interfaces()
	if err != nil {
		panic(err)
	}
	for _, iface := range ifaces {
		if s := iface.HardwareAddr.String(); s != "" {
			return s
		}
	}
	panic("no valid network interfaces")
}

func MustHostname() string {
	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	return hostname
}
