package main

import (
    "fmt"
    "math"
)


type (
	text string
)

func main() {
	var (
		a int
		b float32
		c string
		d [10]int
		e bool
		f byte
		g text
	)

	fmt.Println(a, "int")
	fmt.Println(b, "float32")
	fmt.Println(c, "string")
	fmt.Println(d, "array")
	fmt.Println(e, "bool")
	fmt.Println(f, "byte")
	fmt.Println(g, "self defined of string")

	fmt.Println("\n-- min and max of int8:\n", math.MinInt8)
	fmt.Println(math.MaxInt8, "\n")

	// define and assign
	var h int
	h = 1
	var i float32 = 3.14
	var j = 2.78
	k := 99
	l := false

	fmt.Println(h)
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)
	fmt.Println(l)

	m, _, o, p := 9, 8, 7, 6
	// _ 主要用于返回值的取值，忽略掉其中的某项

	fmt.Println(m)
	fmt.Println(o)
	fmt.Println(p)	 

	// 类型转换
	var q float32 = 100.1
	fmt.Println(q)

	r := uint8(q)
	fmt.Println(r)
}
