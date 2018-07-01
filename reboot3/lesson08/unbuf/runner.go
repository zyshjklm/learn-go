package unbuf

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Runner to run
func Runner() {
	var wg sync.WaitGroup
	ch := make(chan int, 0)

	wg.Add(1)
	fmt.Println("start runner()")
	go running(ch, &wg)
	ch <- 1

	wg.Wait()
}

func running(ch chan int, wg *sync.WaitGroup) {
	var newRunner int

	runner := <-ch
	fmt.Printf("runner %d running with Baton\n", runner)
	if runner != 4 {
		newRunner = runner + 1
		fmt.Printf("runner %d to the line\n", runner)
		go running(ch, wg)
	}

	// rand sleep time
	n := rand.Intn(3)

	time.Sleep(time.Second * time.Duration(n))
	if runner == 4 {
		fmt.Printf("runner %d finish. Race over!\n", runner)
		wg.Done()
		return
	}
	fmt.Printf("runner %d use %d seconds\n", runner, n)
	fmt.Printf("runner %d exchange with runner %d\n\n", runner, runner+1)
	ch <- newRunner
}
