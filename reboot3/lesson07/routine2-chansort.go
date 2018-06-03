package main

import (
	"fmt"
	"sync"
	"time"
)


func worker(i int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Duration(i) * time.Second)
	fmt.Println(i)
}

func main() {
	var ch = make(chan int)
	// 先启动对CHAN的读操作。
	go func() {
		for c := range ch {
			fmt.Println(c)
		}
	}()

	var wg sync.WaitGroup

	s := [5]int{1, 3, 2, 5, 4}
	for _, v := range s {
		wg.Add(1)
		go worker(v, ch, &wg)
	}
	wg.Wait()
}
