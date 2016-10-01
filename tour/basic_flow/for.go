package main

import "fmt"

func main() {
	sum := 0

	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	sum = 1
	// leave the pre and post statement
	for ; sum < 1000 ; {
		sum += sum
	}
	fmt.Println(sum)

	sum = 1
	// drop the ";", C's while is spelled for in Go
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}