package main

import (
	"fmt"
	"log"
	"net/http"
)

// 给定一个url，打印url，和其抓取状态
// http://www.baidu.com 200 OK
func printURL(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Print(err)
		return
	}
	defer resp.Body.Close()
	fmt.Println(url, resp.Status)
}

func work(ch chan string) {
	// work协程从chan中获取url，调用printURL打印
	url := <-ch
	printURL(url)
}

func main() {
	ch := make(chan string)
	urls := "http://www.baidu.com"

	// 主协程启动一个work协程，同时传递一个chan
	// 主协程向chan里发送url
	go work(ch)
	ch <- urls
	// printURL()
}
