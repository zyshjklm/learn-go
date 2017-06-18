package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	var (
		x1 int
		x2 int32
		x3 int64
		x4 uint
		x5 uint32
		x6 uint64
	)
	fmt.Println(x1, x2, x3, x4, x5, x6)
	fmt.Println(reflect.TypeOf(x1), "\t", unsafe.Sizeof(x1))
	fmt.Println(reflect.TypeOf(x2), "\t", unsafe.Sizeof(x2))
	fmt.Println(reflect.TypeOf(x3), "\t", unsafe.Sizeof(x3))
	fmt.Println(reflect.TypeOf(x4), "\t", unsafe.Sizeof(x4))
	fmt.Println(reflect.TypeOf(x5), "\t", unsafe.Sizeof(x5))
	fmt.Println(reflect.TypeOf(x6), "\t", unsafe.Sizeof(x6))

	var (
		x81 int8
		x82 uint8
	)
	fmt.Println(x81, x82)
	fmt.Println(reflect.TypeOf(x81), "\t", unsafe.Sizeof(x81))
	fmt.Println(reflect.TypeOf(x82), "\t", unsafe.Sizeof(x82))

	x81, x82 = 127, 255
	fmt.Println(x81, x82)
	x81 += 1
	x82 += 1
	fmt.Println(x81, x82)
}
