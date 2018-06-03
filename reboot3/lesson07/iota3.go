package main

import "fmt"

const (
	a, b, c = iota + 1, iota + 2, iota + 3
	// iota在一个赋值语句；iota只使用一次。
	// 后面每行的形式和前面一样。
	d, e, f
	g, h, i
)

func main() {
	fmt.Println(a, b, c)
	fmt.Println(d, e, f)
	fmt.Println(g, h, i)
}

/*

go run iota3.go
1 2 3
2 3 4
3 4 5
*/
