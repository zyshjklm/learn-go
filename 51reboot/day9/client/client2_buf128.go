package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	addr := "www.baidu.com:80"
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
	buf := make([]byte, 128)
	for {
		n, err = conn.Read(buf)
		// EOF 代表对方关闭了连接
		if err != nil || err == io.EOF {
			log.Fatal(err)
		}
		fmt.Printf("%s", string(buf[:n]))
	}
	conn.Close()
}
