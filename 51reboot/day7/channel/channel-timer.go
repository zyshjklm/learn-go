package main

import (
	"fmt"
	"time"
)

func main() {

	// after 用于定时，到达时间后，向通道中写入，下条读取通道的语句才能从阻塞中返回
	// 并取得数据
	fmt.Println("start after...")
	c := time.After(time.Second * 2)
	result := <-c
	fmt.Println("done", result)

	timer := time.NewTicker(time.Millisecond * 200)
	cnt := 0
	// if no return inside the for-if.
	// the for loop with encounter such fatal:
	// fatal error: all goroutines are asleep - deadlock!
	for _ = range timer.C {
		cnt++
		if cnt > 5 {
			timer.Stop()
			return
		}
		fmt.Println("hello", cnt)
	}
}
