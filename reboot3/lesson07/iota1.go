package main

import "fmt"

const (
	a = iota
	b // 1
	_ // ignore 2
	c // 3
	d = 15
	e
	f
	g = iota
)

func main() {
	fmt.Println(a, b, c, d, e, f, g)
}

/*
没有显式的赋值，就递增。
go run iota1.go
0 1 3 15 15 15 7
*/
