package main

import "fmt"

func main() {
	fmt.Println(3 / 2)
	fmt.Println(3 / 2.0)
	fmt.Println(3.0 / 2)

	var a, b int
	a, b = 10, 3
	fmt.Println(a, b)
	a = a << 2
	b = b << 3
	fmt.Println(a, b)

	var x uint8
	x = 0xae
	fmt.Printf("%d, %8b\n", x, x)
	x &= 0x7f
	fmt.Printf("%d, %8b\n", x, x)
}
