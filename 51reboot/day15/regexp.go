package main

import (
	"fmt"
	"regexp"
)

var (
	reg = regexp.MustCompile("[0-9]{2,10}")
)

func main() {
	ok := reg.MatchString("abc12")
	fmt.Println(ok)

	fmt.Println(reg.FindString("abc12345"))
}
