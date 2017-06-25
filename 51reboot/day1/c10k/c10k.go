package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func handler(conn net.Conn) {
	fmt.Fprintf(conn, "%s\n", time.Now().String())
	conn.Close()
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handler(conn)
	}
}
