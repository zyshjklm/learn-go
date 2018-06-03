package main

import "fmt"

func main() {
	var ch = make(chan int, 0)

	go func() {
		ch <- 1
	}()

	c := <-ch
	fmt.Println(c)

	// style 2
	var ch2 = make(chan int, 2)
	ch2 <- 1
	ch2 <- 2

	go func() {
		ch2 <- 3
	}()

	fmt.Println(<-ch2)
	fmt.Println(<-ch2)
	fmt.Println(<-ch2)
}
