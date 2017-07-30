package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	addr := "127.0.0.1:8021"
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	buf := make([]byte, 4096)
	rd := bufio.NewReader(os.Stdin)

	line, _ := rd.ReadString('\n')
	line = strings.TrimSpace(line)
	line = fmt.Sprintf("%s HTTP/1.1\r\n\r\n", line)
	n, err := conn.Write([]byte(line))
	if err != nil {
		log.Println("err: ", err)
	}
	log.Println("write size:", n)

	// read
	n, err = conn.Read(buf)
	// EOF 代表对方关闭了连接
	if err != nil || err == io.EOF {
		log.Print("err: ", err)
	}
	log.Printf("return content:\n")
	fmt.Print(string(buf[:n]))
}
