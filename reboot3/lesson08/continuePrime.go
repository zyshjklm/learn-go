package main

import (
	"fmt"
	"runtime"
	"sync"
)

func prime(prefix string, wg *sync.WaitGroup) {
	defer wg.Done()
next:
	for out := 2; out < 10; out++ {
		for in := 2; in < out; in++ {
			fmt.Printf("\t continue:out=%d,in=%d\n", out, in)
			if out%in == 0 {
				continue next
			}
		}
		fmt.Printf("go %s: %d\n", prefix, out)
	}
	fmt.Printf("%s finished\n", prefix)
}

func main() {
	var wg sync.WaitGroup

	runtime.GOMAXPROCS(1)
	wg.Add(2)
	go prime("A", &wg)
	go prime("B", &wg)
	wg.Wait()
}
