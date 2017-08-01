package main

import (
	"io"
	"log"
	"net"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("usage: %s file-for-put", os.Args[0])
	}
	fileName := os.Args[1]
	fd, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer fd.Close()

	addr := "127.0.0.1:8021"
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	buf := make([]byte, 1024)
	_, err = conn.Write([]byte("STORE " + filepath.Base(fileName) + "\n"))
	if err != nil {
		log.Println("err: ", err)
	}

	// write length
	// stat, _ := fd.Stat()
	// length := stat.Size()
	// n, err = conn.Write([]byte(strconv.FormatInt(length, 10)))
	// if err != nil || int64(n) != length {
	// 	log.Fatal("write length error")
	// }
	// log.Println("write size:", n)

	num, err := io.Copy(conn, fd)
	if err != nil || err == io.EOF {
		log.Print("err: ", err)
	}
	log.Print("write num:", num)

	// 发送端使用Copy发送完内容后，服务端使用for循环，
	// 看上去并没有收到EOF。因此使用断方，关闭发送
	// 以便再次发送EOF
	value, _ := conn.(*net.TCPConn)
	value.CloseWrite()
	log.Print("after close write")

	// read
	n, err := conn.Read(buf)
	if err != nil || err == io.EOF {
		log.Print("err: ", err)
	}
	log.Printf("return content:%s", string(buf[:n]))
}
