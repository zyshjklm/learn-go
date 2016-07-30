package main 

import (
	"fmt"
)

func pingPong0(s []int) {
	s = append(s, 3)
	// append后，s是新的地址，所以没有影响到main
}

func pingPong1(s []int) []int {
	s = append(s, 3)
	// append后，s是新的地址，
	return s 
}


func main() {
	
	// about pits of slice
	s0 := make( []int, 0)
	fmt.Println(s0)	// []

	pingPong0(s0)
	fmt.Println(s0)	// []

	// modify
	s1 := make( []int, 0)
	fmt.Println(s1)	// []

	// accept new addr
	s1 = pingPong1(s1)
	fmt.Println(s1)	// [3]

}