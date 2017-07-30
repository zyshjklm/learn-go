package main

import (
	"bufio"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

func server(listener net.Listener) {
	connCh := make(chan net.Conn)
	go worker(connCh)
	go worker(connCh)
	go worker(connCh)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
		}
		connCh <- conn
	}
}

// 读取文件名，打开文件，读取内容，发送内容，关闭连接和文件
func worker(ch chan net.Conn) {
	for {
		conn := <-ch
		rd := bufio.NewReader(conn)
		line, err := rd.ReadString('\n')
		args := strings.Fields(strings.TrimSpace(line))
		if len(args) < 2 || args[0] != "GET" {
			conn.Write([]byte("bad request!\n"))
			continue
		}

		log.Println("get file: ", args[1])
		fd, err := os.Open(args[1])
		if err != nil {
			log.Println(err)
		}
		defer fd.Close()

		time.Sleep(1000 * time.Millisecond)
		log.Print(time.Now().String())

		// 使用Copy虽然可以剩下几行代码，但有个问题就是Copy等待conn先关闭，导致终端hung住
		// io.Copy(conn, fd)
		buf, err := ioutil.ReadAll(fd)
		if err != nil {
			log.Print(err)
		}
		conn.Write(buf)
		conn.Close()
	}
}

func main() {
	addr := ":8021"
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Println(err)
	}
	defer listener.Close()

	server(listener)
}
