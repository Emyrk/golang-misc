package main

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"net"
	"strings"
	"time"
)

var dialer = net.Dialer{
	Timeout:   30 * time.Second,
	KeepAlive: 30 * time.Second,
}

type Message struct {
	Message string
}

func main() {

	fmt.Println("Launching server...")

	// listen on all interfaces
	ln, _ := net.Listen("tcp", ":8081")

	// accept connection on port
	i := 0
	for {
		conn, _ := ln.Accept()
		fmt.Printf("NewConnection %d with remote %s and local %s\n", i, conn.RemoteAddr().String(), conn.LocalAddr().String())
		go dealWithConn(conn, i)
		i++
	}
}

func dealWithConn(conn net.Conn, i int) {
	for {
		enc := gob.NewEncoder(conn)

		//conn.SetWriteDeadline(time.Now().Add(time.Duration(1) * time.Second))
		// run loop forever (or until ctrl-c)
		for {
			// will listen for message to process ending in newline (\n)
			message, err := bufio.NewReader(conn).ReadString('\n')
			if err != nil {
				fmt.Printf("%d has closed: %s\n", i, err.Error())
				return
			}
			// output message received
			fmt.Printf("Message Received from %d:%s\n", i, string(message))
			// sample process for string received
			newmessage := strings.ToUpper(message)
			// send new string back to client
			//conn.Write([]byte(newmessage + "\n"))
			m := new(Message)
			m.Message = newmessage
			fmt.Printf("Writing to %d...\n", i)
			conn.Write([]byte{})
			enc.Encode(m)
		}
	}
}
