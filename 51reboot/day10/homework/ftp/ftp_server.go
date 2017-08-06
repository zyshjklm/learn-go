package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"os"
	"path/filepath"
	"strings"
)

func handleConn(conn net.Conn) {
	// old read method
	/*
		buf := make([]byte, 1024)
		n, _ := conn.Read(buf)
	*/
	// new read method
	r := bufio.NewReader(conn)
	line, _ := r.ReadString('\n')
	fields := strings.Fields(strings.TrimSpace(line))
	if len(fields) != 2 {
		conn.Write([]byte("bad cmd"))
	}
	cmd, name := fields[0], fields[1]
	if cmd == "GET" {
		f, err := os.Open(name)
		if err != nil {
			log.Print(err)
			return
		}
		defer f.Close()
		log.Print("start to copy...")
		io.Copy(conn, f)
		log.Print("after copy...")
	} else if cmd == "STORE" {
		os.MkdirAll(filepath.Dir(name), 0755)
		f, err := os.Create(name)
		if err != nil {
			log.Print(err)
			return
		}
		defer f.Close()
		io.Copy(f, r)
	}
	log.Print("start to close()...")
	conn.Close()
}

func main() {

	addr := ":8021"
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
		}
		log.Print("new connection.\n")
		go handleConn(conn)
	}
}
