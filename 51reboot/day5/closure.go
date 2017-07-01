package main

import "fmt"

// clourse return a function.
type intFunc func(int) int

func addN(n int) intFunc {
	return func(m int) int {
		return m + n
	}
}
func addn(n int) func(int) int {
	return func(m int) int {
		return m + n
	}
}

func main() {
	fn := addn(3)
	fmt.Println(fn(2))

	fN := addN(13)
	fmt.Println(fN(12))
}
