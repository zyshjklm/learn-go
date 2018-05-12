package main

import "fmt"

func main() {
	a := [3]int{1, 2, 3}

	s1 := a[:2]
	s2 := a[:2]
	fmt.Print(s1)
	fmt.Print(s2)

	// fmt.Println(s1 == s2)
	// invalid operation: s1 == s2 (slice can only be compared to nil)
}
