package main

import (
	"context"
	"log"
	"net"
	"sync"

	"github.com/jungle85gopy/learn-go/51reboot/day15/protoBuf/rpcproto3"

	"google.golang.org/grpc"
)

type addrBookStoreServer struct {
	id    int32
	book  *rpcproto.AddrBook
	mutex sync.Mutex
}

// 没有考虑启动时加载数据
func newAddrBookStoreServer() *addrBookStoreServer {
	return &addrBookStoreServer{
		id:   0,
		book: new(rpcproto.AddrBook),
	}
}

//  grpc使用的"golang.org/x/net/context"
func (s *addrBookStoreServer) AddPerson(ctx context.Context, req *rpcproto.AddPersonRequest) (*rpcproto.AddPersonResponse, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	log.Printf("add call:[person:%+v], [phones:%+v]\n", req.Person, req.Person.Phones)

	s.id++
	req.GetPerson().Id = s.id
	// 没有处理去重
	s.book.People = append(s.book.People, req.GetPerson())

	return &rpcproto.AddPersonResponse{
		Id: s.id,
	}, nil
}

func main() {
	ls, err := net.Listen("tcp", ":8021")
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()
	rpcproto.RegisterAddrBookStoreServer(server, newAddrBookStoreServer())
	server.Serve(ls)
}
