package main

import "fmt"

const (
	cPI = 3.1415926
	cE  = 2.0
	cG  = 9.8
)

const (
	cA = iota
	cB
	cC
)

const (
	cRED = iota
	cGREEN
	cBLUE
)
const (
	_ = 1 << (10 * iota)
	cKB
	cMB
	cGB
)

func main() {
	var n int
	// n = PI
	// error| constant 3.14159 truncated to integer
	n = 5
	fmt.Println(cPI, n)

	fmt.Println(cA, cB, cC)
	fmt.Println(cRED, cGREEN, cBLUE)
	fmt.Println(cKB, cMB, cGB)
}
