package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"
)

func TimeoutMiddleware(timeout time.Duration, next func(w http.ResponseWriter, req *http.Request)) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		ctx, cancel := context.WithTimeout(req.Context(), timeout)
		defer cancel()

		req = req.WithContext(ctx)

		next(w, req)

		if err := ctx.Err(); err != nil {
			if errors.Is(err, context.DeadlineExceeded) {
				log.Println("HTTP Request timed out")
				w.Write([]byte("Timed out"))
			}
		}
	}
}

func endless(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	for {
		// This block until the parent ctx is done.
		// The parent ctx is done if the client closes the connection, or the
		// middleware cancels the ctx
		select {
		case <-ctx.Done():
			return
		default:
			fmt.Println("wait")
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", TimeoutMiddleware(time.Second, endless))

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	// http.Client{
	// 	Timeout: time.Second,
	// }.Do

	log.Fatal(server.ListenAndServe())
}
