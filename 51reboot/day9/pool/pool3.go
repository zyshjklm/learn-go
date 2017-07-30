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

	for {
		url, ok := <-ch
		// fmt.Println(url, ok)
		// http://qq.com true
		if !ok {
			break
		}
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

	for i := 0; i < urlNum; i++ {
		go work(ch)
	}
	for i := 0; i < urlNum; i++ {
		ch <- urls[i]
	}
	time.Sleep(2 * time.Second)
	close(ch)
}
