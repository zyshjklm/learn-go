package main

import "fmt"

type gopher struct {
	name string
	age  int
}

func main() {
	var g = gopher{}
	// g.name, age := "smile", 20
	// non-name g.name on left side of :=

	g.name, g.age = "smile", 20
	fmt.Println(g)
}
