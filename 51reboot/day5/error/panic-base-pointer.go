package main

import "fmt"

func main() {
	var pointer *int
	fmt.Println(*pointer)

	// panic: runtime error: invalid memory address or nil pointer dereference
	// [signal SIGSEGV: segmentation violation code=0x1 addr=0x0 pc=0x1093491]
}
