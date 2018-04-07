package cmd

import (
	"log"

	"github.com/jungle85gopy/learn-go/51reboot/day15/protoBuf/rpcproto3"

	grpc "google.golang.org/grpc"
)

// 用于每次命令，每个命令都要调用一个新的client
func newClient(addr string) rpcproto.AddrBookStoreClient {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	return rpcproto.NewAddrBookStoreClient(conn)
}
