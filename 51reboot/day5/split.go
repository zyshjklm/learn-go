package main

import "fmt"

// 命名返回值
func split(sum int) (x, y int) {
	x = sum / 10
	y = sum % 10
	return
}

func main() {
	fmt.Println(split(17))
}
