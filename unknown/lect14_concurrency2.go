package main 

import (
	"fmt"
	"runtime"
	"time"
	"sync"
)


func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	c := make(chan bool)

	for i := 0; i < 10; i++ {
		go testGo(c, i)
	}
	<- c
	// 上述方式会导致运行结果不全

	// 修改方式一 channel length
	time.Sleep(time.Second)
	fmt.Println("\n--- style 1 ---")

	c1 := make(chan bool, 10)
	for i := 0; i < 10; i++ {
		go testGo1(c1, i)
	}

	for i := 0; i < 10; i++ {
		<- c1
	}
	
	// 修改方式二 sync

	fmt.Println("\n--- style 2 ---")
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go testGo2(&wg, i)
	}

	wg.Wait()
}

func testGo(c chan bool, index int) {
	a := 1
	for i := 0; i < 1000000; i++ {
		a += i
	}
	fmt.Println(index, a)

	if index == 9 {
		c <- true
	}
}


func testGo1(c chan bool, index int) {
	a := 1
	for i := 0; i < 1000000; i++ {
		a += i
	}
	fmt.Println(index, a)

	c <- true
}

func testGo2(wg *sync.WaitGroup, index int) {
	a := 1
	for i := 0; i < 1000000; i++ {
		a += i
	}
	fmt.Println(index, a)

	wg.Done()
}
