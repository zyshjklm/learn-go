package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	for i := 0; i < 3; i++ {
		// read from ch1, write to ch2
		go func() {
			x := <-ch1
			ch2 <- x
		}()
	}

	ch1 <- 1
	ch1 <- 2
	ch1 <- 3

	go func() {
		for val := range ch2 {
			fmt.Println(val)
		}
	}()
	time.Sleep(time.Millisecond * 100)
}
