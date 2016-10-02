package main 

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	fmt.Println(strings.Fields(s))

	for _, v := range strings.Fields(s) {
		m[v]++
	}
	return m
}

func main() {
	testString := "hello world by jungle! hello jungle by jie"
	fmt.Println(WordCount(testString))
}