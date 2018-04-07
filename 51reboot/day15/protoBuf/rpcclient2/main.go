package main

import (
	"context"
	"log"

	"github.com/jungle85gopy/learn-go/51reboot/day15/protoBuf/rpcproto2"

	grpc "google.golang.org/grpc"
)

func main() {
	// use http2
	conn, err := grpc.Dial("127.0.0.1:8021", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	var phones []*rpcproto.PhoneNumber
	phone := &rpcproto.PhoneNumber{
		Number: "13812345678",
		Type:   rpcproto.PhoneType_MOBILE,
	}
	phones = append(phones, phone)
	req := &rpcproto.AddPersonRequest{
		Person: &rpcproto.Person{
			Id:     1,
			Name:   "jungle85",
			Email:  "jungle85@github.com",
			Phones: phones,
		},
	}

	client := rpcproto.NewAddrBookStoreClient(conn)
	resp, err := client.AddPerson(context.TODO(), req)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(resp.GetId())
}
