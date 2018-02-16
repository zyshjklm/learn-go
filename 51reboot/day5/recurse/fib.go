package main

import "fmt"

func fib(n int) int {
	if n == 1 || n == 2 {
		return 1
	}

	sum := fib(n-1) + fib(n-2)
	// fmt.Println(sum)
	return sum
}
func main() {
	for i := 1; i < 50; i++ {
		fmt.Println(i, fib(i))
	}
}
