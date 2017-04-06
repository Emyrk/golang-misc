package main

import (
	"bufio"
	"crypto/rand"
	"encoding/gob"
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

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
	random := make([]byte, 10)
	rand.Read(random)
	randomID = fmt.Sprintf("%x", random)
	fmt.Println("I am ", randomID)

	var err error
	// connect to this socket
	conn, err = net.Dial("tcp", "127.0.0.1:8081") //dialer.Dial("tcp", "127.0.0.1:8081")
	for err != nil {
		conn, err = net.Dial("tcp", "127.0.0.1:8081")
	}

	go alwaysRead()
	for {
		// read in input from stdin
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")
		text, _ := reader.ReadString('\n')
		// send to socket
		fmt.Fprintf(conn, text+"\n")
	}
}

func alwaysRead() {
	dec := gob.NewDecoder(conn)
	for {
		//conn.SetReadDeadline(time.Now().Add(time.Duration(1) * time.Second))
		var m Message
		// fmt.Println("Reading...")
		err := dec.Decode(&m) //bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			if err == io.EOF {
				time.Sleep(1 * time.Second)
				conn, err = net.Dial("tcp", "127.0.0.1:8081")
				fmt.Println(err)
				if err == nil {
					dec = gob.NewDecoder(conn)
					fmt.Fprintf(conn, randomID+"\n")
				}
			}
			continue
		}
		fmt.Print("\nMessage from server: " + m.Message)
	}
}
