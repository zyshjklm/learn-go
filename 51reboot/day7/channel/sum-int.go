package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	s := []int{7, 2, 8, 9, 4, 0, 3, 6}

	c := make(chan int)
	half := len(s) / 2

	// 将求和分成2部分进行
	// 两个协程共用一个通道，各用一次来返回求和结果，然后即结束
	go sum(s[:half], c)
	go sum(s[half:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
}
