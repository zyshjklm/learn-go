package main

import "fmt"

func main() {
	var sint [3]int
	idx := 2
	fmt.Println(sint[idx])
	idx++
	fmt.Println(sint[idx])
	// panic: runtime error: index out of range
}
