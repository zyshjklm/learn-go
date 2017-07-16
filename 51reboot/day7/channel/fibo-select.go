package main

import "fmt"

// 两个通道：
// 	quit: 读取，用于控制退出，
//  c   : 写入，用于输出每一轮的结果
func fibonacci(c, quit chan int) {
	x, y := 0, 1

	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)

	// 主动权在主函数，由主控制协程的内部循环运行的次数。
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	//域名函数和对fibo的调用，两者只能一个显式调用协程
	fibonacci(c, quit)

	// 第二方式，
	fmt.Println("-- second --")
	go fibonacci(c, quit)

	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	quit <- 0

}
