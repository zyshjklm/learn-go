package main

import "fmt"

type Point struct {
	X, Y float64
}

func main() {
	var i interface{}

	var n int
	i = n
	fmt.Println("int:", i)

	var s string
	i = s
	fmt.Println("str:", i)

	var p Point
	i = p
	fmt.Println("Point:", i)
}
