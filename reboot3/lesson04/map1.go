package main

import "fmt"

func main() {
	// style 1
	// var ages map[string]int
	// 只声明，却没有初始化
	// panic: assignment to entry in nil map

	// style 2
	// ages := map[string]int{}

	// style 3
	ages := make(map[string]int)

	ages["tome"] = 23
	ages["jack"] = 24

	for k, v := range ages {
		fmt.Printf("%s -> %d\n", k, v)
	}
	delete(ages, "tome")
	fmt.Println("after delete:")

	for k, v := range ages {
		fmt.Printf("%s -> %d\n", k, v)
	}
}
