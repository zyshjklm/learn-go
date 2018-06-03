package main

import "fmt"

func main() {
	s := "hello大家好"

	fmt.Println("len:", len(s), s)

	for i := 0; i < len(s); i++ {
		fmt.Println(string(s[i]))
	}

	r := []rune(s)
	fmt.Println("len:", len(r), r)
	for i := 0; i < len(r); i++ {
		fmt.Println(string(r[i]))
	}
}
