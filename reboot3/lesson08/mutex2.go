package main

import (
	"fmt"
	"strconv"
	"sync"
)

var (
	wg    sync.WaitGroup
	myMap = make(map[string]int64)
)

func opReadMap() {
	defer wg.Done()
	for key, val := range myMap {
		fmt.Println(key, val)
	}
}

func opWriteMap() {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		myMap[strconv.Itoa(i)] = int64(i)
	}
}

func main() {
	wg.Add(2)
	go opReadMap()
	go opWriteMap()

	wg.Wait()
	fmt.Println("over")
}
