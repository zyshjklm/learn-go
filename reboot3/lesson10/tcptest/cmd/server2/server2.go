package server2

import (
	"bytes"
	"encoding/binary"
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
		// Handle the connection in a new goroutine.
		// The loop then returns to accepting, so that
		// multiple connections may be served concurrently.
		go handleNewFunc(conn)
	}
}

// reader process and consume the message form chan
func reader(rdCh chan []byte) {
	for {
		select {
		case msg := <-rdCh:
			fmt.Println(string(msg))
		}
	}
}

func bytes2Int(b []byte) int {
	bb := bytes.NewBuffer(b)
	var x int32
	binary.Read(bb, binary.BigEndian, &x)
	return int(x)
}

func unPacket(btbuf []byte, rdCh chan []byte) []byte {
	length := len(btbuf)
	var i int
	for i = 0; i < length; i++ {
		if length < i+4 {
			break
		}
		msgLen := bytes2Int(btbuf[i : i+4])
		data := btbuf[i+4 : i+4+msgLen]
		rdCh <- data
		i += msgLen + 4 - 1 // from 0
	}
	if i == length {
		return make([]byte, 0)
	}
	return btbuf[i:]
}

func handleNewFunc(c net.Conn) {
	defer c.Close()

	var buf = make([]byte, 1024)
	var tmpBuf = make([]byte, 0)

	readerChan := make(chan []byte, 20)
	go reader(readerChan)

	for {
		n, err := c.Read(buf)
		if err != nil {
			log.Println(err)
			return
		}
		tmpBuf = unPacket(append(tmpBuf, buf[:n]...), readerChan)
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
