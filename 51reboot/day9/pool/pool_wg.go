package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

//
func printURLStatus(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Print(err)
		return
	}
	defer resp.Body.Close()
	fmt.Println(url, resp.Status)
}

func work(ch chan string, wg *sync.WaitGroup) {
	// wg的完成是在协程中表示当前协程的结束。
	defer wg.Done()

	for url := range ch {
		printURLStatus(url)
	}
}

func main() {
	var ch = make(chan string)
	var wg sync.WaitGroup
	// 	wg内部是计数器，Add加数，Done减数。Wait在大于0时阻塞。

	urls := []string{
		"http://www.baidu.com",
		"http://daily.zhihu.com",
		"http://qq.com"}
	urlNum := len(urls)

	// wg.Add(urlNum)
	// Add应该在启动协程前启动。
	// 添加wg的2种方式。上面在循环中，每次加1. 下面是一次完成。
	for i := 0; i < urlNum; i++ {
		wg.Add(1)
		go work(ch, &wg)
	}

	for i := 0; i < urlNum; i++ {
		ch <- urls[i]
	}
	close(ch)
	wg.Wait()
}
