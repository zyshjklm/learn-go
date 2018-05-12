package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var a [3]int = [3]int{1, 2, 3}
	fmt.Println(a[0])
	fmt.Println(a[len(a)-1])

	var a2 = [3]int{1, 3, 5}
	fmt.Println(a2[0])
	fmt.Println(a2[len(a2)-1])

	var r = [5]int{4, 5, 6}
	fmt.Printf("r: %v, len: %d, size:%v\n", r, len(r), unsafe.Sizeof(r))

	b := [...]int{1, 3, 5}
	b[2] = 9
	fmt.Printf("b: %v, len: %d, size:%v\n", b, len(b), unsafe.Sizeof(b))

	c := [...]int{0: 1, 2: 8}
	fmt.Printf("c: %v, len: %d, size:%v\n", c, len(c), unsafe.Sizeof(c))

}
