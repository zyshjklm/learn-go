package main

import (
	"fmt"
	"time"
)

func main() {
	// 如下2条tick赋值语句是等效的。
	// tick := time.NewTicker(1000 * time.Millisecond).C
	tick := time.Tick(1000 * time.Millisecond)
	boom := time.After(4000 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("...滴答...")
		case <-boom:
			fmt.Println("嘣！！！")
			return
		default:
			fmt.Println("吃一口面")
			time.Sleep(400 * time.Millisecond)
		}
	}

}
