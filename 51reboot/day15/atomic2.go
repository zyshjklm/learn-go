package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	cnt := 0
	for {
		var n int32

		wg := new(sync.WaitGroup)
		var lock sync.Mutex
		wg.Add(2)

		go func() {
			lock.Lock()
			defer lock.Unlock()
			atomic.AddInt32(&n, 2)
			// n = n + 2
			wg.Done()
		}()
		go func() {
			lock.Lock()
			defer lock.Unlock()
			n = n / 2
			wg.Done()
		}()
		wg.Wait()
		cnt++
		if n != 2 && n != 1 {
			fmt.Printf("counter:%d, %d\n", cnt, n)
		}
	}
}
