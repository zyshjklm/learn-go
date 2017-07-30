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
	go work(ch)
	go work(ch)
	go work(ch)

	ch <- "http://www.baidu.com"
	ch <- "http://daily.zhihu.com"
	ch <- "http://qq.com"
	time.Sleep(1 * time.Second)
	close(ch)
}
