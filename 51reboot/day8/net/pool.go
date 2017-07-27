package main

import (
	"log"
	"net/http"
	"sync"
	"time"
)

func work(ch chan string, wg *sync.WaitGroup) {
	// read from channel. sleep while no data;
	for u := range ch {
		resp, err := http.Get(u)
		if err != nil {
			log.Print(err)
			return
		}
		time.Sleep(time.Millisecond * 100)
		log.Printf("url: %s, status: %d, len: %d\n", u, resp.StatusCode, resp.ContentLength)
		resp.Body.Close()
	}
	// while channel closed for loop exit.
	wg.Done()
}

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(5)

	taskch := make(chan string)
	for i := 0; i < 5; i++ {
		// start 5 routine
		go work(taskch, wg)
	}

	urls := []string{
		"http://www.baidu.com",
		"http://www.qq.com",
		"http://59.110.12.72:7070/golang-spider/img.html",
		"http://daily.zhihu.com",
	}
	for _, url := range urls {
		taskch <- url
	}
	close(taskch)
	wg.Wait()
}
