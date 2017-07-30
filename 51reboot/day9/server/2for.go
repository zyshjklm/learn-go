package main

import (
	"log"
	"net"
	"time"
)

func server(listener net.Listener) {
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		conn.Write([]byte(time.Now().String() + ": hello golang\n"))
		conn.Close()
	}
}

func main() {
	addr := ":8021"
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	server(listener)
}
