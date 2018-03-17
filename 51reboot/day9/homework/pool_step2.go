package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// 给定一个url，打印url，和其抓取状态
// http://www.baidu.com 200 OK
func printURL(url string) {
	fmt.Println("print url:", url)
	resp, err := http.Get(url)
	if err != nil {
		log.Print(err)
		return
	}
	defer resp.Body.Close()
	fmt.Println(url, resp.Status)
}

func work1(ch chan string) {
	// work协程从chan中获取url，调用printURL打印
	for {
		url, ok := <-ch
		if !ok {
			break
		}
		fmt.Println("work url:", url)
		printURL(url)
	}
}

func work(ch chan string) {
	// work协程从chan中获取url，调用printURL打印
	for url := range ch {
		fmt.Println("work url:", url)
		printURL(url)
	}
}

func main() {
	ch := make(chan string)
	// routine
	for i := 0; i < 3; i++ {
		go work(ch)
	}
	urls := "http://www.zhihu.com"

	for i := 0; i < 6; i++ {
		ch <- urls
	}
	close(ch)
	time.Sleep(time.Second * 3)
}
