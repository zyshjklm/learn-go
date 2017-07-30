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

	urls := []string{
		"http://www.baidu.com",
		"http://daily.zhihu.com",
		"http://qq.com"}
	urlNum := len(urls)

	for i := 0; i < urlNum; i++ {
		wg.Add(1)
		go work(ch, &wg)
	}

	for i := 0; i < urlNum; i++ {
		ch <- urls[i]
	}
	close(ch)

	// 启动不要用Sleep来进行协程的同步。一定会出问题。前面只是演示
	// time.Sleep(2 * time.Second)
	wg.Wait()
}
