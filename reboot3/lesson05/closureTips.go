package main

import "fmt"

func test1() {
	var funcList []func()

	for i := 0; i < 3; i++ {
		funcList = append(funcList, func() {
			// 这里循环时，只是在修改切片，参数i
			fmt.Println(i)
		})
	}

	for _, f := range funcList {
		f()
	}
}

func test2() {
	var funcList []func()

	for i := 0; i < 3; i++ {
		var x = i
		funcList = append(funcList, func() {
			fmt.Println(x)
		})
	}

	for _, f := range funcList {
		f()
	}
}
func main() {
	test1()
	test2()
}

/*
go run closureTips.go
3
3
3
*/
