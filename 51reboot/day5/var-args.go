package main

import (
	"fmt"
	"reflect"
)

// 如果定义时使用...，表示省略
// args实际是一个切片来实现的
func sum(args ...int) int {
	n := 0
	fmt.Println("type of args:", reflect.TypeOf(args))
	for i := 0; i < len(args); i++ {
		n += args[i]
	}
	return n
}

func main() {
	fmt.Println(sum(1, 2, 3, 4, 5))
	s := []int{1, 3, 5, 7, 9}
	// 调用时使用...，表示自动展开切片的元素
	fmt.Println(sum(s...))
}
