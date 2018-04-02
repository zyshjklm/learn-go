package main

import (
	"fmt"
	"unsafe"
)

type StringHeader struct {
	Data unsafe.Pointer
	Len  int
}

func slice(s string, b int, len int) string {
	hdr := *(*StringHeader)(unsafe.Pointer(&s))
	// hdr是原string结构的地址，是一个整体
	hdr.Data = unsafe.Pointer(uintptr(hdr.Data) + uintptr(b))
	hdr.Len = len
	s1 := *(*string)(unsafe.Pointer(&hdr))
	return s1
}

func ZeroCopyString(buf []byte) string {
	hdr := &StringHeader{
		Data: unsafe.Pointer(&buf[0]),
		Len:  len(buf),
	}
	return *(*string)(unsafe.Pointer(hdr))
}

func main() {
	str := "abcdef"
	s5 := slice(str, 2, 2)
	fmt.Println(s5)

	// buf := []byte{'h', 'e', 'l', 'l', 'o'}
	buf := []byte("golang")
	str2 := string(buf)
	// string() 有一次值拷贝，两者的值不同
	buf[0] = 'X'
	fmt.Println(string(buf), str2)

	// 没有值拷贝过程，两者的值相同
	str3 := ZeroCopyString(buf)
	fmt.Println(string(buf), str3)
}
