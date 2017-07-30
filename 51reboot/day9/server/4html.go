package main

import (
	"log"
	"net"
	"time"
)

var content = `HTTP/1.1 200 OK
Date: Sat, 29 Jul 2017 06:18:23 GMT
Content-Type: text/html
Connection: Keep-Alive
Server: BWS/1.1
X-UA-Compatible: IE=Edge,chrome=1
BDPAGETYPE: 3
Set-Cookie: BDSVRTM=0; path=/

<html>
<body>
<h1 style="color:red">hello golang</h1>
</body>
</html>
`

func worker(ch chan net.Conn) {
	for {
		conn := <-ch
		time.Sleep(1000 * time.Millisecond)
		conn.Write([]byte(content))
		conn.Close()
	}
}

func server(listener net.Listener) {
	connCh := make(chan net.Conn)
	go worker(connCh)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
		}
		connCh <- conn
	}
}

func main() {
	addr := ":8021"
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	server(listener)
}
