package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

var (
	wg       sync.WaitGroup
	myRWLock sync.RWMutex
	myMap    = make(map[string]int64)
)

func opReadMap(id int) {
	defer wg.Done()

	time.Sleep(time.Second * 1)
	// read lock
	myRWLock.RLock()
	for key, val := range myMap {
		fmt.Printf("-- read id=%d: key=%s, val=%d\n", id, key, val)
	}
	myRWLock.RUnlock()
}

func opWriteMap() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Println("== write ==", i)
		// Lock locks rw for writing.
		myRWLock.Lock()
		myMap[strconv.Itoa(i)] = int64(i)
		myRWLock.Unlock()
		time.Sleep(time.Millisecond * 200)
	}
}

func main() {
	wg.Add(1)

	go opWriteMap()
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go opReadMap(i)
	}

	wg.Wait()
	fmt.Println("over")
}
