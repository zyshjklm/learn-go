package main

import "fmt"

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	c := make(chan int, 10)

	// 第一个参数控制了协程内部循环运行的次数
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}
