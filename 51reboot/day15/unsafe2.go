package main

import (
	"fmt"
	"unsafe"
)

type T struct {
	m int8
	n int64
}

func main() {
	var n int
	fmt.Println(unsafe.Sizeof(n)) //8

	var t T
	fmt.Println(unsafe.Sizeof(t)) // 16
	// 结构体存在对齐问题，m是1个字节，但占8字节。
	fmt.Println(unsafe.Alignof(t.m)) // 1
	fmt.Println(unsafe.Alignof(t.n)) // 8

	// 某个字段在struct中的偏移量
	fmt.Println(unsafe.Offsetof(t.m)) // 0
	fmt.Println(unsafe.Offsetof(t.n)) // 8

}
