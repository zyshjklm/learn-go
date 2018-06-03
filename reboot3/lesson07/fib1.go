package main

import (
	"flag"
	"fmt"
)

var (
	num int
)

func init() {
	flag.IntVar(&num, "n", 0, "number")
	flag.Parse()
}

func fib1(n int) int {
	if n < 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return n
	}
	return fib1(n-1) + fib1(n-2)
}

func fib2(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i <= n; i++ {
		//fmt.Println("iter:", i)
		ch <- y
		x, y = y, x+y
	}
	close(ch)
}

func main() {
	var ch = make(chan int)
	fmt.Println("num:", num)
	fmt.Println("fib1 :", fib1(num))

	go fib2(num, ch)

	/*
		for i:=0; i< num;i++{
			fmt.Println(<-ch)
		}
	*/
	for c := range ch {
		fmt.Println("fib2 :", c)
	}
}
