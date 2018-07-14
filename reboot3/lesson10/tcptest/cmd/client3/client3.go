package client3

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"sync"

	"github.com/jkak/learn-go/reboot3/lesson10/tcptest/cmd/common"
)

// ConnectServer connect to server
func ConnectServer(host string, port uint16) {
	addr, err := common.NewServerInfo("tcp", host, port)
	if err != nil {
		log.Fatal(err)
	}

	conn, err := net.Dial(addr.Proto, fmt.Sprintf("%s:%d", addr.Host, addr.Port))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	common.SetTimeout(conn, 3)
	var wg sync.WaitGroup
	wg.Add(1)
	go sender(conn, &wg)
	wg.Wait()
}

func int2byte(n int) []byte {
	x := int32(n)
	bb := bytes.NewBuffer([]byte{})
	binary.Write(bb, binary.BigEndian, x)
	return bb.Bytes()
}

// Packet to packet msg
func Packet(msg []byte) []byte {
	s := make([]byte, 0)
	return append(append(s, int2byte(len(msg))...), msg...)
}

func sender(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 11; i++ {
		data := fmt.Sprintf("{\"ID\":%d, \"Name\":\"user-%d\"}", i, i)
		n, err := conn.Write(Packet([]byte(data)))
		if err != nil {
			log.Println(err.Error())
			continue
		}
		fmt.Printf("index %d: write %d data\n", i, n)
	}
}
