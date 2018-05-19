package main

import "fmt"

type Student struct {
	ID   int
	Name string
	Sex  bool
}

func main() {
	var s Student
	s.ID = 1
	s.Name = "jack"

	p := &s
	fmt.Println(p)
	p.ID = 10
	fmt.Println(p)

	s1 := Student{ID: 2, Name: "alice"}
	fmt.Println(s1)
}
