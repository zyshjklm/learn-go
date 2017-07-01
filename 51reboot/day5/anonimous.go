package main

import (
	"fmt"
	"strings"
)

// strings.Map
// func Map(mapping func(rune) rune, s string) string
// http://localhost:6060/pkg/strings/#Map

func toupper1(s string) string {
	return strings.Map(func(r rune) rune {
		return r - ('a' - 'A')
	}, s)
}

func toupper2(s string) string {
	return strings.Map(smap, s)
}
func smap(r rune) rune {
	fmt.Printf("smap func: %c\n", r)
	return r - ('a' - 'A')
}

func main() {
	fmt.Println(toupper1("hello"))
	fmt.Println(toupper2("golang"))

}
