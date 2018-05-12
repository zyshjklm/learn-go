package main

import "fmt"

func main() {
	a := make([]int, 5)
	prints("a", a)

	b := make([]int, 0, 5)
	prints("b", b)

	c := b[:2]
	prints("c", c)

	d := c[2:5]
	prints("d", d)

	d = append(d, 2)
	prints("d", d)

	d = append(d, 4, 5, 6)
	prints("d", d)
}

func prints(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(s), cap(x), x)
}
