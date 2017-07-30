package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	addr := "www.baidu.com:80"
	// connection 流式数据
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	// write
	n, err := conn.Write([]byte("GET / HTTP/1.1\r\n\r\n"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("\nwrite size:", n, "\nreturn content:\n")

	// read
	rd := bufio.NewReader(conn)
	for {
		line, err := rd.ReadString('\n')
		if err != nil || err == io.EOF {
			log.Fatal(err)
		}
		fmt.Printf("%s", line)
	}
	conn.Close()
}
