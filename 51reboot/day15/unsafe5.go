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
	s := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(&s[0])

	hdr2 := SliceHeader{
		Data: unsafe.Pointer(&s[0]),
		Len:  3,
		Cap:  6,
	}

	// s2 指向了s
	s2 := *(*[]int)(unsafe.Pointer(&hdr2))
	fmt.Println(len(s2))
	fmt.Println(cap(s2))
	fmt.Println(s2[0], s2[1])
	fmt.Println("end hdr2 test --:\n")
	//
	hdr3 := SliceHeader{
		Data: unsafe.Pointer(&s),
		Len:  3,
		Cap:  6,
	}

	s3 := *(*[]int)(unsafe.Pointer(&hdr3))
	fmt.Println(len(s3))
	fmt.Println(cap(s3))
	fmt.Println(s3[0], s3[1])
	fmt.Println("end hdr3 test --:\n")

	// test
	fmt.Println(s)
	s4 := slice(s, 1, 2)
	fmt.Println(len(s4))
	fmt.Println(cap(s4))
	fmt.Println(s4[0], s4[1])
}

func slice(s []int, b int, len int) []int {
	hdr := SliceHeader{
		Data: unsafe.Pointer(&s[b]),
		Len:  int64(len),
		Cap:  int64(cap(s)),
	}
	s1 := *(*[]int)(unsafe.Pointer(&hdr))
	return s1
}
