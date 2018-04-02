package main

import (
	"fmt"
	"unsafe"
)

type SliceHeader struct {
	Data unsafe.Pointer
	Len  int64
	Cap  int64
}

func main() {
	// data pointer
	// len int64 长度
	// cap int64 总容量

	s := []int{1, 2, 3}
	fmt.Println(&s[0])

	var p *SliceHeader
	p = (*SliceHeader)(unsafe.Pointer(&s))
	fmt.Printf("%#v\n", *p)

	s1 := s[:1]
	p = (*SliceHeader)(unsafe.Pointer(&s1))
	fmt.Printf("%#v\n", *p)
	// s, s1, 与两次p中的地址是一个位置
}
