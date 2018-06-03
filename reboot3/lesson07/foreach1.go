package main

import "fmt"

type gopher struct {
	name string
	age  int
}

func main() {
	var l = []gopher{}
	var g = gopher{}
	g.name, g.age = "smile", 20

	l = append(l, g)
	fmt.Println(g)
	for _, r := range l {
		fmt.Println(r)
	}

	my := []gopher{
		gopher{"jk", 30},
	}
	fmt.Println(my)
	my1 := []gopher{
		{name: "tom", age: 11},
		{name: "jk", age: 11},
		{name: "alice", age: 11},
	}
	fmt.Println(my1)
}
