package main

import "fmt"

var g = "smile"

// syntax error: mixed named and unnamed function parameters
// func test1() (c1 string, error) {
// 	return "s1", nil
// }

// 不推荐强命名
func test1() (c1 string, err error) {
	return "s1", nil
}

func test2() (string, error) {
	return "s1", nil
}

func main() {
	fmt.Println(test1())
	fmt.Println(test2())
}
