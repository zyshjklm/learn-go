package main

import (
	"fmt"
	"strings"
)

func toUpper(s string) string {
	mapFunc := func(r rune) rune {
		return r - ('a' - 'A')
	}
	return strings.Map(mapFunc, s)
}

func toUpper2(s string) string {
	mapper := func(r rune) rune {
		return r - 32
	}
	return strings.Map(mapper, s)
}

func main() {
	fmt.Println(toUpper("hello"))
	fmt.Println(toUpper2("hello"))
}
