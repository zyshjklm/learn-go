package main 

import (
	"fmt"
	"strconv"
)



func main() {
	var x int = 65
	y := string(x)
	z := strconv.Itoa(x)		// int to str
	x1, _ := strconv.Atoi(z)	// str to int

	fmt.Println(x, "int")		// 65
	fmt.Println(y, "string")	// A 	[ not "65" ]
	fmt.Println(z, "string")	// "65"
	fmt.Println(x1, "int\n")	// to int


	const (
		a1 = "123"
		b1 = len(a1)
		c1
	)

	fmt.Println(a1)
	fmt.Println(b1)
	fmt.Println(c1, "\n")

	const (
		a2, b2 = 1, "22"
		c2, d2		// used by couple.
	)

	fmt.Println(a2)
	fmt.Println(b2)
	fmt.Println(c2)
	fmt.Println(d2, "\n")

	// iota
	const (
		a3 = 'A'	// 65
		b3 			// 65， 使用上行的常量表达式
		c3 = iota 	// 2， 从第一个常量开始递增
		d3 			// 3
	)

	fmt.Println(a3)
	fmt.Println(b3)
	fmt.Println(c3)
	fmt.Println(d3, "\n")

	const (
		e = 'A'		// 65
		f = iota	// 1
		g 		 	// 2
		h 			// 3
	)	

	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)

	const (
		i = 'A'		// 65
		j = iota	// 1
		k = 'B' 	// 66
		l = iota	// 3 
	)	

	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)
	fmt.Println(l)
}

