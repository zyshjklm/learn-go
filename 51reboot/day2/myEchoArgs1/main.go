package main

import (
	"flag"
	"fmt"
	"os"
)

var newline = flag.Bool("n", true, "print new line")

func main() {
	var s string
	fmt.Println(os.Args)
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		// sep from sep.go
		sep = " "
	}
	fmt.Println(*newline)

	if *newline {
		fmt.Print("arg line:")
		fmt.Println(s)
	} else {
		fmt.Print("arg noline:")
		fmt.Print(s)
	}
}
