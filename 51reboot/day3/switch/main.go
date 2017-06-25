package main

import "fmt"

func main() {
	s := "1"
	switch s {
	case "1":
		fmt.Println("s1=", "1")
	case "2":
		fmt.Println("s1=", "2")
	default:
		fmt.Println("s1=", "default", s)
	}
}
