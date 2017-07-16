package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Student1 struct {
	Name string
	id   int
}
type Student2 Student1

func (s Student2) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.id)
}

func main() {
	s1 := Student1{
		Name: "binggan",
		id:   1,
	}
	s2 := Student2(s1)
	buf1, err := json.Marshal(s1)
	if err != nil {
		log.Fatal(err)
	}

	buf2, err := json.Marshal(s2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(buf1))
	// {"Name":"binggan"}
	// 上面未实现序列化方法，与下面实现之后的差别
	fmt.Println(string(buf2))
	// 1
}
