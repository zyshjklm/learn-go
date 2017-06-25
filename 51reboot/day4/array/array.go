package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var q = [3]int{1, 2, 3}
	var r = [3]int{1, 2}
	fmt.Println(q[2])
	fmt.Println(r[2])

	q1 := [...]int{1, 2, 3, 4}
	fmt.Println(q1)

	q2 := [...]int{4: 2, 10: -1}
	fmt.Println(q2)

	a1 := [3]int{1, 2, 3}
	var a2 [3]int

	a2 = a1
	fmt.Println(&a1[0], &a2[0])    //地址不相同，发生了值拷贝
	fmt.Println(unsafe.Sizeof(a1)) // 24
	fmt.Println(unsafe.Sizeof(q1)) // 32
	fmt.Println(unsafe.Sizeof(q2)) // 88
}
