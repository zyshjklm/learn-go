package main

import (
	"fmt"
	"io"
	"net"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1)
	listener, err := net.Listen("tcp", "127.0.0.1:8848")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		go handle(conn.(*net.TCPConn))
	}
}

func handle(server *net.TCPConn) {
	defer server.Close()
	client, err := net.Dial("tcp", "test.com:8849")
	if err != nil {
		fmt.Print(err)
		return
	}
	defer client.Close()
	go func() {
		defer server.Close()
		defer client.Close()
		buf := make([]byte, 2048)
		io.CopyBuffer(server, client, buf)
	}()
	//  io.Copy 的默认 buffer 比较大，给一个小的 buffer 可以支持更多的并发连接
	buf := make([]byte, 2048)
	io.CopyBuffer(client, server, buf)
}
