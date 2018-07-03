package main

import (
	"fmt"
	"unsafe"

	"github.com/jkak/learn-go/reboot3/lesson09/pkgbase/meta"
)

func main() {
	m := meta.New()
	m.Y = 100
	fmt.Println(m)

	// m.x = 10
	// m.x undefined (cannot refer to unexported field or method x)
	// fmt.Println(m)

	p := (*struct{ z int32 })(unsafe.Pointer(m))
	p.z = 5
	// 这时可以设置原本不可见的x变量
	fmt.Println(m)

	q := (*struct {
		m int
		n int
	})(unsafe.Pointer(m))
	q.m = 6
	q.n = 9
	// n的类型不对，且原本是可见的。现在却不可见了
	fmt.Println(m)

	r := (*struct {
		m int32
		n int32
	})(unsafe.Pointer(m))
	r.m = 16
	r.n = 19
	fmt.Println(m)
}
