package main

import (
	"fmt"
	"runtime/debug"
)

func main() {
	recoverCall()
}

func recoverCall() {
	defer func() {
		fmt.Println("in defer: stack info:")
		if err := recover(); err != nil {
			debug.PrintStack()
			fmt.Println("recover!")
		}
		fmt.Println("in defer: stack info:")
	}()
	call()
}

func call() {
	defer func() { fmt.Println("1") }()
	defer func() { fmt.Println("2") }()
	defer func() { fmt.Println("3") }()

	panic("panic error")
}
