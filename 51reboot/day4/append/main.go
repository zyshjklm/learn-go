package main

import "fmt"

func main() {
	var s = make([]int, 0, 1)
	fmt.Println(s)

	for i := 1; i < 18; i++ {
		s = append(s, i)
		printSlice(s)
	}
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
