package main

import (
	"fmt"
)

type Student struct {
	Id   int
	Name string
}

func main() {
	var s Student
	s.Id = 1
	s.Name = "jack where are you going!"

	// output
	fmt.Println(s)
	fmt.Printf("%v\n", 1)
	fmt.Println(fmt.Sprintf("http://%s/%s", "www.baidu.com", "about.html"))
	fmt.Println("http://%s/%s", "www.baidu.com", "about.html")

	// input
	var str string
	var i int
	fmt.Print("input a string and int: ")
	fmt.Scanln(&str, &i)
	fmt.Println(str, i)
	s1 := Student{
		Id:   2,
		Name: "alice",
	}
	fmt.Println(s1)

	s1 = s
	fmt.Println(s1)

}
