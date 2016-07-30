package main 

import (
	"fmt"
	"time"
)


func main() {
	c1, c2 := make(chan int), make(chan string)
	o := make(chan bool, 2)

	// multi channel
	go func() {
		for {
			select {	// 多个可用时，选取是随机的
			case v, ok := <- c1:
				if !ok {
					// fmt.Println("-- c1")
					o <- true
					break
				}
				fmt.Println("c1", v)

			case v, ok := <- c2:
				if !ok {
					// fmt.Println("-- c2")
					o <- true
					break
				}
				fmt.Println("c2", v)
			}
		}
	}()

	c2 <- "hi"	
	c1 <- 1


	c1 <- 3
	c2 <- "hello"

	// close(c1)
	close(c2)

	for i := 0; i < 2; i++ {
		<- o
	}

	// write and read channel by select 
	c := make(chan int)
	go func() {
		for v := range c {
			fmt.Println(v)
		}
	}()

	for i := 0; i < 10; i++ {
		select {
			case c <- 0:
			case c <- 1:
		}
	}

	// timeout of select
	c3 := make(chan bool)
	select {
		case v := <- c3:
			fmt.Println(v)
		case <- time.After(2 * time.Second):
			fmt.Println("Timeout")

	}
}


