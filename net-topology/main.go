package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
	"time"
)

type Connection struct {
	Connected  bool   `json:"Connected"`
	Hash       string `json:"Hash"`
	Connection struct {
		MomentConnected  time.Time `json:"MomentConnected"`
		BytesSent        int       `json:"BytesSent"`
		BytesReceived    int       `json:"BytesReceived"`
		MessagesSent     int       `json:"MessagesSent"`
		MessagesReceived int       `json:"MessagesReceived"`
		PeerAddress      string    `json:"PeerAddress"`
		PeerQuality      int       `json:"PeerQuality"`
		ConnectionState  string    `json:"ConnectionState"`
		ConnectionNotes  string    `json:"ConnectionNotes"`
	} `json:"Connection"`
	ConnectionTimeFormatted string `json:"ConnectionTimeFormatted"`
	PeerHash                string `json:"PeerHash"`
}

func main() {
	var (
		start = flag.String("s", "localhost", "Control panel to start with")
	)
	flag.Parse()

	n := NewNetwork()
	n.ExpandNode(Node{*start, 8090, false})

	network := n.String()

	file, err := os.OpenFile("network.txt", os.O_CREATE|os.O_WRONLY, 0777)
	if err == nil {
		file.WriteString(network)
		file.Close()
	}
	fmt.Println(len(n.Nodes))
	fmt.Println(count)
	fmt.Printf("Average Connection Count %d\n", totalConnections/len(n.Nodes))

}

type Node struct {
	Loc      string
	Port     int
	Searched bool
}

func (n Node) String() string {
	return fmt.Sprintf("%s:%d", n.Loc, n.Port)
}

type Network struct {
	Connections []string
	Nodes       map[string]Node

	sync.Mutex
}

func NewNetwork() *Network {
	n := new(Network)
	n.Nodes = make(map[string]Node)
	return n
}

func (n *Network) String() string {
	str := ""
	for _, s := range n.Connections {
		str += s + "\n"
	}
	return str
}

var count int
var totalConnections int

func (n *Network) ExpandNode(s Node) {
	n.Lock()
	if f := n.Nodes[s.String()]; f.Searched {
		n.Unlock()
		return
	}

	fmt.Printf("Searching %s...\n", s.String())

	// Search
	s.Searched = true
	n.Nodes[s.String()] = s
	n.Unlock()

	cons, err := getConnections(s.Loc, s.Port)
	if err != nil {
		return
	}

	totalConnections += len(cons)
	for i, c := range cons {
		o := Node{c.Connection.PeerAddress, 8090, false}
		n.ExpandNode(o)
		if i <= 4 {
			//fmt.Println(i)
			count++
			n.Connections = append(n.Connections, fmt.Sprintf("%s -- %s", s.String(), o.String()))
		} else {
			//fmt.Println("WHY")
		}
	}
}

func getConnections(loc string, port int) ([]Connection, error) {
	client := http.Client{Timeout: 3 * time.Second}

	resp, err := client.Get(fmt.Sprintf("http://%s:%d/factomd?item=peers", loc, port))
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var connections []Connection
	err = json.Unmarshal(data, &connections)
	if err != nil {
		return nil, err
	}

	return connections, nil
}
