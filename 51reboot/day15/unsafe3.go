package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var n int
	var m int64
	var p *int

	p = &n
	*p = 32
	fmt.Printf("%d\n", *p)
	// p = (*int)(&m)
	// ./unsafe3.go:16:12: cannot convert &m (type *int64) to type *int
	m = int64(n)

	// 使用Pointer
	p = (*int)(unsafe.Pointer(&m))
	*p = 0xFFFFFFFFA
	fmt.Println(*p)

	var m2 [2]int8
	fmt.Println(m2)

	p = (*int)(unsafe.Pointer(&m2))
	*p = 0x0101
	fmt.Println(*p)
	// 影响了2个字节
}
