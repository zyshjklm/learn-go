package main

import "fmt"

func printStr(s string) {
	defer func() {
		fmt.Println("defer in str")
	}()
	fmt.Println(s)
}

func printRet(s string) {
	defer func() {
		fmt.Println("defer in ret")
	}()
	return
	fmt.Println(s)
}

func main() {
	printStr("hello")
	printRet("golang")
}
