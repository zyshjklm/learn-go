package main

import (
	"fmt"
	"sync"
)

func worker(msg string, ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	ch <- msg
}

func main() {
	var ch = make(chan string)
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for c := range ch {
			fmt.Println(c)
		}
	}()

	go worker("work", ch, &wg)

	wg.Wait()
}