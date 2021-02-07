package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
)

type HijackAccept struct {
	net.Listener
}

func (j *HijackAccept) Accept() (net.Conn, error) {
	conn, err := j.Listener.Accept()
	if err != nil {
		return conn, err
	}

	host, _, err := net.SplitHostPort(conn.RemoteAddr().String())
	if err == nil {
		if host == "127.0.0.1" {
			conn.Close()
			return nil, fmt.Errorf("blocked %s", host)
		}
	}

	log.Panic()
	log.Fatal()

	return conn, err
}

func main() {
	// l, err := net.Listen("tcp", "127.0.0.1:8090")
	// if err != nil {
	// 	panic(err)
	// }

	// hijackedListener := &HijackAccept{Listener: l}

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		_, _ = fmt.Fprintf(w, "Hello world")
	})

	server := http.Server{
		Addr:    "192.168.1.6:8080",
		Handler: mux,
	}
	server.ListenAndServe()

	// log.Fatal(http.ListenAndServe(":8090", mux))
	// log.Fatal(http.Serve(l, mux))
}
