package main

import "fmt"

func main() {
	s := make([]int, 3)
	s = append(s, 1, 2, 3)
	fmt.Println(s)

	s1 := make([]int, 0)
	s1 = append(s1, 1, 2, 3)
	fmt.Println(s1)
}

/*
go run slice1.go
[0 0 0 1 2 3]
[1 2 3]
*/
