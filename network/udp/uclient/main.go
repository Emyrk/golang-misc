package main

import (
	"context"
	"fmt"
	"io"
	"net"
	"time"
)

// maxBufferSize specifies the size of the buffers that
// are used to temporarily hold data from the UDP packets
// that we receive.
const maxBufferSize = 1024

var timeout = time.Second * 5

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	client(ctx, "localhost:1234", nil)
	select {
	case <-ctx.Done():
	}
}

// client wraps the whole functionality of a UDP client that sends
// a message and waits for a response coming back from the server
// that it initially targetted.
func client(ctx context.Context, address string, reader io.Reader) (err error) {
	// Resolve the UDP address so that we can make use of DialUDP
	// with an actual IP and port instead of a name (in case a
	// hostname is specified).
	raddr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		return
	}

	// Although we're not in a connection-oriented transport,
	// the act of `dialing` is analogous to the act of performing
	// a `connect(2)` syscall for a socket of type SOCK_DGRAM:
	// - it forces the underlying socket to only read and write
	//   to and from a specific remote address.
	conn, err := net.DialUDP("udp", nil, raddr)
	if err != nil {
		return
	}

	// Closes the underlying file descriptor associated with the,
	// socket so that it no longer refers to any file.
	defer conn.Close()

	doneChan := make(chan error, 1)

	_, _ = conn.Write([]byte("hi"))

	go func() {
		time.Sleep(time.Second * 4)
		for {
			// It is possible that this action blocks, although this
			// should only occur in very resource-intensive situations:
			// - when you've filled up the socket buffer and the OS
			//   can't dequeue the queue fast enough.
			//n, err := io.Copy(conn, reader)
			//if err != nil {
			//	doneChan <- err
			//	return
			//}
			//
			//fmt.Printf("packet-written: bytes=%d\n", n)

			buffer := make([]byte, maxBufferSize)

			// Set a deadline for the ReadOperation so that we don't
			// wait forever for a server that might not respond on
			// a resonable amount of time.
			deadline := time.Now().Add(timeout)
			err = conn.SetReadDeadline(deadline)
			if err != nil {
				doneChan <- err
				return
			}

			nRead, addr, err := conn.ReadFromUDP(buffer)
			if err != nil {
				doneChan <- err
				return
			}

			fmt.Printf("packet-received: bytes=%d from=%s :: %v\n",
				nRead, addr.String(), string(buffer[:nRead]))

			//doneChan <- nil
		}
	}()

	select {
	case <-ctx.Done():
		fmt.Println("cancelled")
		err = ctx.Err()
	case err = <-doneChan:
	}

	return
}
