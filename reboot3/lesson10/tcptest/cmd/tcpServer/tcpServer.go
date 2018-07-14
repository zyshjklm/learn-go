package tcpServer

import (
	"fmt"
	"io"
	"log"
	"net"

	"github.com/jkak/learn-go/reboot3/lesson09/tcptest/cmd/common"
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
		go func(c net.Conn) {
			io.Copy(c, c)
			c.Close()
		}(conn)
	}
}
