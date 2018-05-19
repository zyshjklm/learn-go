package main

import "fmt"

func sum(args ...int) int {
	n := 0
	for i := 0; i < len(args); i++ {
		n += args[i]
	}
	return n
}

func minMax(args ...int) (min, max int) {
	min = args[0]
	max = args[0]
	for i := 1; i < len(args); i++ {
		if min > args[i] {
			min = args[i]
		}
		if max < args[i] {
			max = args[i]
		}
	}
	return
}

func main() {
	fmt.Println(sum(1, 2, 3, 4))
	fmt.Println(minMax(1, 2, 3, 4))
}
