package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := "abc123"
	n, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
		// parsing "abc123": invalid syntax
		// return or exit
	} else {
		fmt.Println(n)
	}

	s = "123"
	m, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(m)

	s = "456"
	if x, err := strconv.Atoi(s); err == nil {
		fmt.Println(x)
	}
}
