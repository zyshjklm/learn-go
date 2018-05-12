package main

import "fmt"

func main() {
	s := make([]int, 0, 3)

	for i := 0; i < 5; i++ {
		s = append(s, i)
		fmt.Printf("cap:%d, len:%d, s:%p, value:%v\n", cap(s), len(s), s, s)
	}
}
