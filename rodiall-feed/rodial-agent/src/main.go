package main

import (
	"fmt"
	"time"
	"os"
	"net/http"
)

func main() {
	// Minimal proof-of-life agent: exposes a simple HTTP health endpoint
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "{\"status\":\"ok\",\"pid\":%d}", os.Getpid())
	})
	go http.ListenAndServe(":8081", nil)

	// Simulate a persistent connection to FBIS (placeholder)
	for {
		fmt.Println("rodial-agent: heartbeating to FBIS (placeholder)")
		time.Sleep(30 * time.Second)
	}
}
