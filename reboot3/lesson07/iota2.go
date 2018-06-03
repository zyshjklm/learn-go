package main

import "fmt"

const (
	a, b = iota + 1, iota + 2
	// iota在一个赋值语句；iota只使用一次。
	// 后面每行的形式和前面一样。
	c, d
	e, f
)

func main() {
	fmt.Println(a, b, c, d, e, f)
}

/*

go run iota2.go
1 2 2 3 3 4

*/
