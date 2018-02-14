package main

import "fmt"

func main() {
	var s = make([]int, 0, 1)
	printSlice(s)

	for i := 1; i < 18; i++ {
		s = append(s, i)
		printSlice(s)
	}
}

func printSlice(s []int) {
	fmt.Printf("len=%2d cap=%2d %v\n", len(s), cap(s), s)
}
