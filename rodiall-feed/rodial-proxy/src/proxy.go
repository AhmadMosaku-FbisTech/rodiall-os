package main
import (
	"fmt"
	"net"
	"os"
)

func main() {
	fmt.Println("rodial-proxy: starting placeholder")
	ln, err := net.Listen("tcp", "127.0.0.1:1080")
	if err != nil {
		fmt.Println("listen error", err)
		os.Exit(1)
	}
	for {
		conn, err := ln.Accept()
		if err != nil { continue }
		conn.Close()
	}
}
