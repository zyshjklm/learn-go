package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/gogo/protobuf/proto"
	"github.com/jungle85gopy/learn-go/51reboot/day15/protoBuf/myproto"
)

func main() {
	var p myproto.Person
	buf, _ := ioutil.ReadAll(os.Stdin)
	err := proto.Unmarshal(buf, &p)
	if err != nil {
		panic(err)
	}
	fmt.Println(p)
	fmt.Println(p.String())
}
