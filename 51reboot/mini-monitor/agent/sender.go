package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"

	"github.com/jungle85gopy/learn-go/51reboot/mini-monitor/common"
)

// Sender for agent
type Sender struct {
	addr string
	ch   chan *common.Metric
}

// NewSender 构造函数
func NewSender(addr string) *Sender {
	sender := &Sender{
		addr: addr,
		ch:   make(chan *common.Metric, 1024),
	}
	return sender
}

// Channel 返回ch
func (s *Sender) Channel() chan *common.Metric {
	return s.ch
}

// Start 建立连接；
// 循环从ch中读取metric，序列化metric，发送数据
func (s *Sender) Start() {
	var conn net.Conn
	conn, err := net.Dial("tcp", s.addr)
	if err != nil {
		panic(err)
	}
	for {
		metric := <-s.ch
		buf, _ := json.Marshal(metric)
		_, err := fmt.Fprintf(conn, "%s\n", buf)
		if err != nil {
			log.Printf("send metric err:%s", err.Error())
		}
	}
}
