package main

import (
	"log"
	"net/rpc"

	"github.com/jungle85gopy/learn-go/51reboot/day15/rpc/common"
)

func main() {
	client, err := rpc.Dial("tcp", "127.0.0.1:8021")
	if err != nil {
		log.Fatal(err)
	}
	req := common.AddRequest{
		M: 10,
		N: 16,
	}
	var reply common.AddResponse

	err = client.Call("MathService.Add", &req, &reply)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("result:%d\n", reply.Result)
}
