package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"log"
	"net"

	"github.com/jungle85gopy/learn-go/51reboot/mini-monitor/common"
)

var addr = flag.String("addr", ":6000", "transfer port")

func main() {
	flag.Parse()
	lisn, err := net.Listen("tcp", *addr)
	if err != nil {
		log.Fatal(err)
	}
	defer lisn.Close()

	for {
		conn, err := lisn.Accept()
		if err != nil {
			log.Print(err)
			return
		}
		go handleConn(conn)
	}
}

// 按行读取；反序列化成结构
func handleConn(conn net.Conn) {
	defer conn.Close()

	var metric common.Metric
	r := bufio.NewReader(conn)
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			log.Print(err)
		}
		if len(line) == 0 {
			continue
		}
		// 去掉换行符，有2种常用方法
		// line = strings.TrimSpace(line)
		line = line[:len(line)-1]

		json.Unmarshal([]byte(line), &metric)
		log.Print("metric :", metric)
		// 反序列化有2种方式，一种是定义变量，一种new一个指针
		metric2 := new(common.Metric)
		json.Unmarshal([]byte(line), metric2)
		log.Print("metric2:", *metric2)
	}
}
