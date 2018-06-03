package main

import (
	"fmt"
	"sync"
)

func worker(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("worker")
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go worker(&wg)
	wg.Wait()
}
