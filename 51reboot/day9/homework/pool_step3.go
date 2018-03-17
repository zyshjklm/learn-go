package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
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

func work(ch chan string, wg *sync.WaitGroup) {
	// work协程从chan中获取url，调用printURL打印
	defer wg.Done()
	for url := range ch {
		fmt.Println("work url:", url)
		printURL(url)
	}
}

func main() {
	ch := make(chan string)
	wg := new(sync.WaitGroup)
	// wg.Add(3)
	// routine
	for i := 0; i < 3; i++ {
		wg.Add(1) // 在
		go work(ch, wg)
	}
	urls := "http://www.zhihu.com"

	for i := 0; i < 6; i++ {
		ch <- urls
	}
	close(ch)
	wg.Wait()
}
