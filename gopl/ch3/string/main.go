// practise 3.4 string
package main 

import (
	"fmt"
)


func main() {
	// string
	s := "hello, world!"
	length := len(s)

	fmt.Println(length)
	fmt.Println(s[0], s[7])	// 104 119

	// slice
	fmt.Println(s[:5])
	fmt.Println(s[5:])
	fmt.Println(s[:])

	unicode := string('国')
	length = len(unicode)

	fmt.Println(length)
	for i := 0; i < length; i++ {
		fmt.Println(unicode[i])
	}
	// 229 155 189

	fmt.Println(unicode[:])

	// add 
	fmt.Println("s :", s)
	t := s
	s += " by jungle!"

	fmt.Println("s :", s)
	fmt.Println("t :", t)
	
}