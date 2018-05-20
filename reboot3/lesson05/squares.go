package main

import "fmt"

func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

func main() {
	var f func() int = squares()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())

	var f1 func() int = squares()
	fmt.Println(f1())
	fmt.Println(f1())
	fmt.Println(f1())
}

/*
go run squares.go
1
4
9
16
1
4
9
*/
