package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	counter int64
	wg      sync.WaitGroup
	myLock  sync.Mutex
)

func incCounter(id int) {
	defer wg.Done()

	for i := 0; i < 2; i++ {
		myLock.Lock()
		{
			value := counter
			runtime.Gosched()
			value++
			counter = value
		}
		myLock.Unlock()
	}
}

func main() {
	//runtime.GOMAXPROCS(1)

	wg.Add(2)
	go incCounter(1)
	go incCounter(2)

	wg.Wait()
	fmt.Println("out:", counter)
}
