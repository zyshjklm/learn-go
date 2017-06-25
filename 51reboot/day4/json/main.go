package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Student struct {
	Id   int
	Name string
}

func main() {
	s := Student{
		Id:   1,
		Name: "jack",
	}
	buf, err := json.Marshal(s)
	if err != nil {
		log.Fatalf("marshal err: %s", err)
	}
	str := string(buf)
	fmt.Println(str)

	// un
	var s1 Student
	err = json.Unmarshal([]byte(str), &s1)
	if err != nil {
		log.Fatalf("unmarshal err:%s", err)
	}
	fmt.Printf("%v\n", s1)
	fmt.Printf("%#v\n", s1)

	str = `{"Id":2, "Name":"alice"}`
	err = json.Unmarshal([]byte(str), &s1)
	if err != nil {
		log.Fatalf("unmarshal err:%s", err)
	}
	fmt.Printf("%v\n", s1)
	fmt.Printf("%#v\n", s1)
}
