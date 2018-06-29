package main

import "fmt"

func main() {
	var i int
	fmt.Println("use goto statement to implement for loop:")
FOR:
	if i < 10 {
		fmt.Println(i)
		i++
		goto FOR
	}
	fmt.Println("over")
}
