package main

import (
	"encoding/json"
	"testing"

	"github.com/gogo/protobuf/proto"
	"github.com/jungle85gopy/learn-go/51reboot/day15/protoBuf/myproto"
)

func BenchmarkProto(b *testing.B) {
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

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := proto.Marshal(&p)
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkJSON(b *testing.B) {
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

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := json.Marshal(&p)
		if err != nil {
			panic(err)
		}
	}
}
