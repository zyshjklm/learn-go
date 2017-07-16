package main

import (
	"fmt"
	"time"
)

func main() {
	s := []int{2, 7, 1, 5, 4, 3}
	for _, num := range s {
		go func(n int) {
			time.Sleep(time.Duration(n) * time.Millisecond * 200)
			fmt.Println(n)
		}(num)
	}
	// 这里的时间需要够长，不然主协程退出，则所有协程都退出了。
	time.Sleep(2 * time.Second)
}
