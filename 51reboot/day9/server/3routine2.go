package main

import (
	"log"
	"net"
	"time"
)

func server(listener net.Listener) {
	connCh := make(chan net.Conn)
	go worker(connCh)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		connCh <- conn
	}
}

func worker(ch chan net.Conn) {
	for {
		conn := <-ch
		time.Sleep(1000 * time.Millisecond)
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
