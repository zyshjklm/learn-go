package main

import "fmt"

type Student struct {
	ID   int
	Name string
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

	s1 = s
	fmt.Println(s1)
}

/*
{1 jack}
{2 alice}
{1 jack}
*/
