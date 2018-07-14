package client2

import (
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

func sender(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 30; i++ {
		data := fmt.Sprintf("{\"ID\":%d, \"Name\":\"user-%d\"}", i, i)
		n, err := conn.Write([]byte(data))
		if err != nil {
			log.Println(err.Error())
			return
		}
		fmt.Printf("%d: write %d data\n", i, n)
	}
}
