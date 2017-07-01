package main

import (
	"fmt"
	"os"
)

func myprint() {
	fmt.Println("hello")
}

func main() {
	_, err := os.Open("abcd.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	myprint()
}
