package main 

import "fmt"

// package level var
var c, python, java bool

// variable with initializer.
var j, k = 1, 2



func main() {
	// function level var
	var i int
	var shell, golang = true, "yes"

	// short variable declaration
	x := 5
	fmt.Println(c, python, java)
	fmt.Println(shell, golang)
	fmt.Println(i, j, k, x)
}

