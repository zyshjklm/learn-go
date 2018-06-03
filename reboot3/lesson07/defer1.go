package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("in defer")
		if err := recover(); err != nil {
			fmt.Println("recover!")
		}
	}()

	call()
}

func call() {
	fmt.Println("start call...")
	defer func() { fmt.Println("1") }()
	defer func() { fmt.Println("2") }()
	defer func() { fmt.Println("3") }()

	fmt.Println("start panic...")
	panic("done")
	fmt.Println("never be here...")
}
