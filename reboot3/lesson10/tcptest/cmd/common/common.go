package common

import (
	"fmt"
	"net"
	"time"
)

// ServerInfo server info
type ServerInfo struct {
	Proto string
	Host  string
	Port  uint16
}

// NewServerInfo return ServerINfo
func NewServerInfo(proto, host string, port uint16) (*ServerInfo, error) {
	if len(host) == 0 || len(proto) == 0 || port <= 0 {
		return nil, fmt.Errorf("host, proto or port error")
	}
	return &ServerInfo{
		Host:  host,
		Port:  port,
		Proto: proto,
	}, nil
}

// SetTimeout for timeout
func SetTimeout(conn net.Conn, timeout int) {
	conn.SetReadDeadline(time.Now().Add(time.Duration(timeout) * time.Second))
	conn.SetWriteDeadline(time.Now().Add(time.Duration(timeout) * time.Second))
}
