package main

import "fmt"

// a(n+1) = 2 * a(n) + n
func fib(n int) int {
	if n == 1 {
		return 2
	}
	return 2*fib(n-1) + (n - 1)
}
func main() {
	for i := 1; i < 12; i++ {
		fmt.Println(i, fib(i))
	}
}
