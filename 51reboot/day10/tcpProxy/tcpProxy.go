package main

import (
	"flag"
	"io"
	"log"
	"net"
	"sync"
)

var (
	domain = flag.String("domain", "www.baidu.com:80", "target domain host")
)

func handleConn(conn net.Conn) {
	log.Printf("start to handle conn...\n\n")
	// 建立到目标服务器的连接
	remote, err := net.Dial("tcp", *domain)
	if err != nil {
		log.Print(err)
		conn.Close()
		return
	}
	log.Println("connect to domain:", *domain)
	wg := new(sync.WaitGroup)
	wg.Add(2)

	// 从源client读数据，发送remote；直到conn的EOF，关闭remote
	// 需要使用go，因为两个Copy是阻塞式的。
	go func() {
		defer wg.Done()
		log.Println("go rd start...")
		io.Copy(remote, conn)
		log.Println("go rd end...")
	}()

	// go 接收remote的数据，发送给client
	go func() {
		defer wg.Done()
		log.Println("go wr start...")
		io.Copy(conn, remote)
		log.Println("go wr end...")
	}()

	log.Println("start to wait()...")
	wg.Wait()
	log.Println("end of wait()...")

	remote.Close()
	conn.Close()
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
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			return
		}
		go handleConn(conn)
	}
}
