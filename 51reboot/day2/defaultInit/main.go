package main

import "fmt"

func main() {
	var (
		x int
		y float32
		z string
		p *int
		b bool
	)

	fmt.Printf("%v\n", x)
	fmt.Printf("%v\n", y)
	fmt.Printf("%v\n", z)
	fmt.Printf("%v\n", p)
	fmt.Printf("%v\n", b)
	// 0, 0, "", <nil>, false

	i := 0
	s := "hello"
	i, j := 1, 2 // bad init
	fmt.Println(i, j, s)
	// 1 2 hello
}
