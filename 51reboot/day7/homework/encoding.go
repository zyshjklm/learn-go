package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Student struct {
	Name string
	Id   int
}

// StudentUn. 首字母大小写，影响可见性，也影响json的序列化
type StudentUn struct {
	Name string
	id   int
}

func main() {
	s := Student{
		Name: "binggan",
		Id:   1,
	}
	buf, err := json.Marshal(s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(buf))
	// {"Name":"binggan","Id":1}

	s1 := StudentUn{
		Name: "reboot",
		id:   2,
	}
	buf1, err := json.Marshal(s1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(buf1))
	// {"Name":"reboot"}
	// 小写的id，影响了其对序列化的可见性。
}
