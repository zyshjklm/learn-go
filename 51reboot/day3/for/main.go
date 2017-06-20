package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("-- style 1 --")
	for i := 2; i < 3; i++ {
		fmt.Println(i)
	}

	fmt.Println("-- style 2, like while --")
	j := 5
	for j < 8 {
		fmt.Println(j)
		j++
	}

	fmt.Println("-- style 3, like while true --")
	i := 8
	for {
		i++
		fmt.Println(i)
		if i > 10 {
			break
		}
	}
	fmt.Println("-- style 4, range --")
	for i, arg := range os.Args {
		fmt.Println(i, arg)
	}

	s := "hello"
	for i, c := range s {
		fmt.Println(i, c, string(c))
	}

	// 100以内的斐波那契数之和
	fmt.Println(fib(100))
}

func fib(end int) int {
	a, b := 0, 1
	var result int
	for b < end {
		a, b = b, a+b
		result += a
		fmt.Printf("%d ", a)
	}
	fmt.Println()
	return result
}
