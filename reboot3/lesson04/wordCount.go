package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "one two three one"
	var words []string

	words = strings.Fields(s)
	bufWord := make(map[string]int)
	for _, val := range words {
		bufWord[val] += 1
	}
	for k, v := range bufWord {
		fmt.Printf("%s\t%d\n", k, v)
	}
}
