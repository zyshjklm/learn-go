package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	go func(c chan int) {
		for i := 1; i < 10; i++ {
			c <- i
		}
	}(ch)

	for i := 1; i < 10; i++ {
		fmt.Printf("%d: %d\n", i, <-ch)
	}
	// use channel us buffer
	ch <- 1
	ch <- 2

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
