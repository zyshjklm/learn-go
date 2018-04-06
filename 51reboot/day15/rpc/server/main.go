package main

import (
	"log"
	"net"
	"net/rpc"

	"github.com/jungle85gopy/learn-go/51reboot/day15/rpc/common"
)

// MathService as service
type MathService struct {
}

// Add as method of MathService
func (m *MathService) Add(req *common.AddRequest, reply *common.AddResponse) error {
	log.Printf("call add:%v", req)
	reply.Result = req.M + req.N
	return nil
}

func main() {
	ms := new(MathService)
	rpc.Register(ms)

	l, err := net.Listen("tcp", ":8021")
	if err != nil {
		log.Fatal(err)
	}

	rpc.Accept(l)
}
