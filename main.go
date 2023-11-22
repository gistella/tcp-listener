package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
)

const defaultPort uint16 = 8080

func main() {
	var port uint16

	p, err := strconv.ParseUint(os.Getenv("LISTEN_PORT"), 10, 16)
	if err != nil {
		port = defaultPort
	} else {
		port = uint16(p)
	}

	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	fmt.Printf("Listening on port %d\n", port)

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Print(err)
			os.Exit(1)
		}
		defer conn.Close()
		go handleConnection(conn)
	}
}

func handleConnection(c net.Conn) {
	fmt.Printf("Accepted connection from %s\n", c.RemoteAddr())
}
