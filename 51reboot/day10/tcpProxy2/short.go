package main

import (
	"flag"
	"io"
	"log"
	"net"
)

var domain = flag.String("domain", "www.qq.com:80", "target domain host")

func handleConn(conn net.Conn) {
	log.Printf("start to handle conn...\n")
	remote, err := net.Dial("tcp", *domain)
	if err != nil {
		log.Print(err)
		conn.Close()
		return
	}
	defer conn.Close()
	defer remote.Close()

	go io.Copy(remote, conn)
	io.Copy(conn, remote)

	log.Println("end of handle...")
}

func main() {
	flag.Parse()
	log.Println("domain: ", *domain)

	addr := ":8021"
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		conn, _ := listener.Accept()
		go handleConn(conn)
	}
}
