package main

import "fmt"

func main() {
	a := [2][4]int{
		[4]int{1, 2, 3, 4},
		[4]int{5, 6, 7, 8},
	}
	row := len(a)
	col := len(a[0])
	fmt.Println(a)
	fmt.Printf("row: %d, col: %d\n", row, col)

	b := [2][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}}
	fmt.Println(b)
}
