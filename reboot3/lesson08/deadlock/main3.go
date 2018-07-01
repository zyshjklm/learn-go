package main

import (
	"fmt"
	"sync"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1)
		// read from ch1, write to ch2
		go func() {
			defer wg.Done()
			x := <-ch1
			ch2 <- x
		}()
	}

	ch1 <- 1
	ch1 <- 2
	ch1 <- 3

	go func() {
		wg.Wait()
		close(ch2) // avoid deadlock
	}()

	// 这里不再使用routine
	for val := range ch2 {
		fmt.Println(val)
	}

	//time.Sleep(time.Second * 2)
}
