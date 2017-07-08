package main

import "fmt"

// Student show stu
type Student struct {
	Id   int
	Name string
}

// Student1 show stu1
type Student1 struct {
	Id   int
	Name string
}

func main() {
	m := make(map[string]Student)
	m["bingan"] = Student{
		Id:   1,
		Name: "bingan",
	}
	fmt.Println(m)
	// m["bingan"].Id = 2
	fmt.Println(m)

	m1 := make(map[string]*Student1)
	m1["bingan"] = &Student1{
		Id:   1,
		Name: "bingan",
	}
	fmt.Println(m1)
	m1["bingan"].Id = 2
	fmt.Println(m1)
}
