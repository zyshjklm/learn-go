package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(i int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Duration(i) * 100 * time.Millisecond)
	fmt.Println("-- worker process:", i)
	ch <- i
}

func main() {
	var ch = make(chan int)
	var quit = make(chan bool)
	var wg sync.WaitGroup

	go func() {
		for {
			select {
			case c := <-ch:
				fmt.Println("select:", c)
			case <-quit:
				fmt.Println("\nselect bye")
				return
			default:
				fmt.Println("\tdefault sleep...")
				time.Sleep(50 * time.Millisecond)
			}
		}
	}()

	s := [5]int{1, 3, 2, 5, 4}

	for _, i := range s {
		wg.Add(1)
		fmt.Println("start:", i)
		go worker(i, ch, &wg)
	}
	wg.Wait()
	quit <- true
	time.Sleep(100 * time.Millisecond)
}
