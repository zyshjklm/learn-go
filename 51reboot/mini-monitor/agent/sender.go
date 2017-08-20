package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"time"

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

// connect retry connect to transfer.
func (s *Sender) connect() net.Conn {
	baseGap := 100 * time.Microsecond
	for {
		conn, err := net.Dial("tcp", s.addr)
		if err != nil {
			log.Print(err)
			time.Sleep(baseGap)
			baseGap *= 2
			if baseGap > time.Second*30 {
				baseGap = time.Second * 30
			}
			continue
		}
		log.Printf("local addr:%s\n", conn.LocalAddr())
		return conn
	}
}

// reConnect retry connect while write to conn err
func (s *Sender) reConnect(conn net.Conn) *bufio.Writer {
	conn.Close()
	conn = s.connect()
	w := bufio.NewWriter(conn)
	return w
}

// Start 建立连接；
// 循环从ch中读取metric，序列化metric，发送数据
func (s *Sender) Start() {
	var conn net.Conn
	conn = s.connect()
	w := bufio.NewWriter(conn)

	ticker := time.NewTicker(time.Second * 5)
	for {
		select {
		case metric := <-s.ch:
			buf, _ := json.Marshal(metric)
			_, err := fmt.Fprintf(w, "%s\n", buf)
			if err != nil {
				log.Printf("Fprintf to remote err:%s", err.Error())
				w = s.reConnect(conn)
			}
		case <-ticker.C:
			// log.Print("-- Flush data to transfer from bufio of conn.")
			err := w.Flush()
			if err != nil {
				log.Printf("Flush to remote err:%s", err.Error())
				w = s.reConnect(conn)
			}
		}
	}
}
