package client

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"

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
	fmt.Println("send...")
	fmt.Fprintf(conn, "[%+v] this is a client ...\n", time.Now())

	fmt.Println("recv...")
	cont, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cont)
}
