package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	counter int64
	wg      sync.WaitGroup
)

func incCounter(id int) {
	defer wg.Done()

	for {
		atomic.AddInt64(&counter, 1)
		runtime.Gosched()
	}
}

func main() {
	runtime.GOMAXPROCS(2)

	wg.Add(2)
	go incCounter(1)
	go incCounter(2)

	wg.Wait()
	fmt.Println("out:", counter)
}
