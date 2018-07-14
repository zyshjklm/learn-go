package server3

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/jkak/learn-go/reboot3/lesson10/tcptest/cmd/common"
)

const (
	MaxQueueSize = 5
)

type Job struct {
	conn net.Conn
}

var jobQueue = make(chan Job, MaxQueueSize)

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

	// process job from chan
	go func() {
		for {
			time.Sleep(time.Second * 1)
			select {
			case job := <-jobQueue:
				go handleNewFunc(job.conn)
			}
		}
	}()

	for {
		conn, err := l.Accept()
		fmt.Printf("conn from %+v\n", conn.RemoteAddr().String())
		if err != nil {
			log.Fatal(err)
		}
		// go handleNewFunc(conn)
		if len(jobQueue) == MaxQueueSize {
			fmt.Println("---- job queue full...")
		}
		jobQueue <- Job{conn: conn}
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
		msgLen := bytes2Int(btbuf[i : i+4])
		if length < i+4+msgLen {
			break
		}
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
