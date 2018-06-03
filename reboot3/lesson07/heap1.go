package heap

import (
	"fmt"
)

var g = "smile"

func test1() {
	var a [1]int
	c := a[:]
	fmt.Println(c)
}

func test2() {
	var a [1]int
	c := a[:]
	println(c)
}

func test3() {
	p := &g
	println(p)
}

// go tool compile -S heap.go | more

// call runtime.newobject(SB)
// runtime负责对堆和栈的管理。不需要程序员干预。

/*
go tool compile -m heap1.go
heap1.go:15:6: can inline test2
heap1.go:21:6: can inline test3
heap1.go:12:13: c escapes to heap
heap1.go:11:8: a escapes to heap
heap1.go:10:6: moved to heap: a
heap1.go:12:13: test1 ... argument does not escape
heap1.go:17:8: test2 a does not escape
heap1.go:22:7: test3 &g does not escape
 */

// escape, 逃逸，原本在栈，却放在堆里。
