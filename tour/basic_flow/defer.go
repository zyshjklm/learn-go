package main 

import (
	"fmt"
)

func defer_test() {
	defer fmt.Println("World")
	fmt.Println("Hello")
}

func defer_stack() {
	fmt.Println("\n---- counting ----\n")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done in func")
}

func main() {
	defer_test()
	defer_stack()
	fmt.Println("done in main().")
}