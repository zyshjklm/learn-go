package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(1)
	fmt.Println("Max:", runtime.NumCPU())

	var wg sync.WaitGroup
	wg.Add(3)
	go a2zLower(&wg)
	go a2zUpper(&wg)
	go num(&wg)

	fmt.Println("waiting to finish")
	wg.Wait()
	fmt.Println("Finished")
}

func a2zLower(wg *sync.WaitGroup) {
	defer wg.Done()
	for count := 0; count < 2; count++ {
		for ch := 'a'; ch < 'a'+26; ch++ {
			fmt.Printf("%c ", ch)
		}
	}
	fmt.Println("\nlower finish")
}

func a2zUpper(wg *sync.WaitGroup) {
	defer wg.Done()
	for count := 0; count < 2; count++ {
		for ch := 'A'; ch < 'A'+26; ch++ {
			fmt.Printf("%c ", ch)
		}
	}
	fmt.Println("\nupper finish")
}

func num(wg *sync.WaitGroup) {
	defer wg.Done()
	for count := 0; count < 5; count++ {
		for n := 0; n < 10; n++ {
			fmt.Printf("%d ", n)
		}
	}
	fmt.Println("\nnum finish")
}
