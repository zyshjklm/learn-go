package main

import (
	"fmt"
)

func worker(si []int, ch chan int) {
	var result int
	for _, v := range si {
		result += v
	}
	fmt.Println("worker result:", result)
	ch <- result
}

func main() {
	var ch = make(chan int)

	s := []int{1, 3, 6, 4, 0, -9}
	// 前3个数求和，后3个求和，分别放入CHAN
	halfLen := len(s) / 2
	go worker(s[0:halfLen], ch)
	go worker(s[halfLen:], ch)

	x, y := <-ch, <-ch
	fmt.Println("sum:", x+y)
}
