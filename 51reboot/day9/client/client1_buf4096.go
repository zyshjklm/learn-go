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

	// 远端地址和本地地址。
	fmt.Println(conn.RemoteAddr().String())
	fmt.Println(conn.LocalAddr().String())

	// write
	n, err := conn.Write([]byte("GET / HTTP/1.1\r\n\r\n"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("\nwrite size:", n)

	// read
	buf := make([]byte, 4096)
	n, err = conn.Read(buf)
	// EOF 代表对方关闭了连接
	if err != nil || err == io.EOF {
		log.Fatal(err)
	}
	// buf[:n], 因为长度是4096，但读回的长度可能没有填满整个buf
	// 因此只打印读到的部分。
	fmt.Printf("\nreturn content:\n%s\n", string(buf[:n]))
}
