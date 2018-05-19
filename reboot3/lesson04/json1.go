package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Sex  bool
}

func main() {
	var s Student
	s.ID = 1
	s.Name = "jack"
	fmt.Println(s)

	s1 := Student{
		ID:   2,
		Name: "alice",
	}
	fmt.Println(s1)

	buf, _ := json.Marshal(s)
	fmt.Printf("json buf is:%s\n", string(buf))

	buf2, _ := json.MarshalIndent(s, "", "\t")
	fmt.Printf("json buf is:\n%s\n", string(buf2))
}
