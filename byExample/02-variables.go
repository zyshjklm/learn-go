package main

import "fmt"

func main() {
	var a = "initial"
	fmt.Println(a)

	//declare multiple variables at once.
	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	// var declared without a corresponding initialization are zero-valued.
	var e int
	fmt.Println(e)

	// := syntax is shorthand for declaring and initializing a var
	f := "short"
	fmt.Println(f)
}
