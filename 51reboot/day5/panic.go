package main

import "fmt"

func panicPointer() {
	// var p *int
	// fmt.Println(*p)
	// panic: runtime error: invalid memory address or nil pointer dereference
	// [signal SIGSEGV: segmentation violation code=0x1 addr=0x0 pc=0x1087111]
}

func panicIndex() {
	var slice [3]int
	var i = 2
	// when i >= 3:
	// panic: runtime error: index out of range
	fmt.Println(slice[i])
}

func panicZero() {
	var n int
	n = 1
	// n = 0 by default.
	// panic: runtime error: integer divide by zero
	fmt.Println(10 / n)
}

func main() {
	panicPointer()
	panicIndex()
	panicZero()

	defer func() {
		err := recover()
		fmt.Println(err)
	}()

	panic("不想执行下去了")

}
