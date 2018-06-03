package main

import "fmt"

type gopher struct {
	name string
	age  int
}

func main() {
	var m = make(map[string]*gopher)

	my := []gopher{
		{name: "tom", age: 11},
		{name: "jk", age: 12},
		{name: "alice", age: 13},
	}

	for i, r := range my {
		// modify name of tom
		if r.name == "tom" {
			my[i].name = "wang"
			my[i].age += 10
		}
		m[r.name] = &my[i]
	}

	for k, v := range m {
		fmt.Println(k, v)
	}
}

/*
tom &{wang 21}
jk &{jk 12}
alice &{alice 13}
*/
