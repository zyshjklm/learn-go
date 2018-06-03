package main

import (
	"fmt"
	"sync"
	"time"
)

var result []int

func sortList(i int,  wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Duration(i) * time.Second)
	result = append(result, i)
	fmt.Println("sorted:", i)
}

func main() {
	var wg sync.WaitGroup

	s := [5]int{1, 3, 2, 5, 4}
	for _, v := range s {
		wg.Add(1)
		go sortList(v, &wg)
	}
	wg.Wait()

	for _, v := range result {
		fmt.Println(v)
	}
}
