package main

import (
	"fmt"
	"unsafe"
)

type Student struct {
	Id   int
	Name string
}

func main() {
	var s Student
	s.Id = 1
	s.Name = "jack where are you going to fly! "
	var str = "str"
	fmt.Println(unsafe.Sizeof(s))      // 24
	fmt.Println(unsafe.Sizeof(s.Id))   // 8
	fmt.Println(unsafe.Sizeof(s.Name)) //
	fmt.Println("\ninfo of str:")
	fmt.Println(unsafe.Sizeof(str)) // 16
	fmt.Println(len(str))           // 3

	var x int8
	fmt.Println(unsafe.Sizeof(x))

	// style 2 for init
	s1 := Student{
		Id:   2,
		Name: "alice",
	}
	fmt.Println(s1)

	s1 = s
	fmt.Println(s1)

	// struct pointer
	var p *Student
	p = &s1
	p.Id = 3
	fmt.Println(s1)

	var p1 *int
	p1 = &s1.Id
	*p1 = 4
	fmt.Println(s1)
}
