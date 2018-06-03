package main

import (
	"fmt"
	"time"
)

func worker() {
	fmt.Println("worker")
}

func main() {
	go worker()
	 time.Sleep(1 * time.Second)
	 // 如果不等待一下，则worker协程来不及执行，主程序就退出了，看不到结果
	 // 显然，Sleep只能用于示范问题，不能用于线上做协程同步
}
