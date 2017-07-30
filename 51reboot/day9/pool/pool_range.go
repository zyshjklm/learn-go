package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func printURLStatus(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Print(err)
		return
	}
	defer resp.Body.Close()
	fmt.Println(url, resp.Status)
}

func work(ch chan string) {
	for url := range ch {
		printURLStatus(url)
	}
	return
}

func main() {
	var ch = make(chan string)
	urls := []string{
		"http://www.baidu.com",
		"http://daily.zhihu.com",
		"http://qq.com"}
	urlNum := len(urls)

	for i := 0; i < urlNum+2; i++ {
		go work(ch)
	}
	for i := 0; i < urlNum; i++ {
		ch <- urls[i]
	}
	// 启动不要用Sleep来进行协程的同步。一定会出问题。前面只是演示
	time.Sleep(2 * time.Second)
	close(ch)
}
