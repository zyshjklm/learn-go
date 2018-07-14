package server2

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/jkak/learn-go/reboot3/lesson10/tcptest/cmd/common"
)

// TCPServer for tcp server
func TCPServer(host string, port uint16) {
	addr, err := common.NewServerInfo("tcp", host, port)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("listenning...")
	l, err := net.Listen(addr.Proto, fmt.Sprintf("%s:%d", addr.Host, addr.Port))
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		fmt.Printf("conn from %+v\n", conn.RemoteAddr().String())
		if err != nil {
			log.Fatal(err)
		}
		go handleNewFunc(conn)
	}
}

func handleNewFunc(c net.Conn) {
	defer c.Close()

	var buf = make([]byte, 1024)
	for {
		n, err := c.Read(buf)
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Println(n, string(buf[:n]))
	}
}

func handleFunc(c net.Conn) {
	defer c.Close()
	time.Sleep(10 * time.Second)

	var buff = make([]byte, 1024)
	n, err := c.Read(buff)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(buff))

	n, err = c.Write(buff)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Printf("write %d data\n", n)
}
