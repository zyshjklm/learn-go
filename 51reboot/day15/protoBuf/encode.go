package main

import (
	"fmt"

	"github.com/gogo/protobuf/proto"
	"github.com/jungle85gopy/learn-go/51reboot/day15/protoBuf/myproto"
)

func main() {
	var p myproto.Person
	p.Id = 1
	p.Name = "jungle"
	p.Email = "jungle@github.com"
	p.Phones = []*myproto.PhoneNumber{
		{
			Number: "18612349876",
			Type:   myproto.PhoneType_MOBILE,
		},
		{
			Number: "87651234",
			Type:   myproto.PhoneType_HOME,
		},
	}
	buf, err := proto.Marshal(&p)
	if err != nil {
		panic(err)
	}
	fmt.Print(string(buf))
	// fmt.Print(buf)
}
