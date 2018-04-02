package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var n int
	fmt.Println(unsafe.Sizeof(n))
	// 不同的平台，值不一样
}
