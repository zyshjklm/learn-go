package main

import (
	"fmt"
	"strconv"
	"sync"
)

var (
	wg     sync.WaitGroup
	myLock sync.Mutex
	myMap  = make(map[string]int64)
)

func opReadMap() {
	defer wg.Done()
	for key, val := range myMap {
		fmt.Println(key, val)
	}
}

func opWriteMap() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		myLock.Lock()
		myMap[strconv.Itoa(i)] = int64(i)
		myLock.Unlock()
	}
}

func main() {
	wg.Add(2)

	// 是否可能读不到数据？
	go opReadMap()
	go opWriteMap()

	wg.Wait()
	fmt.Println("over")
}
