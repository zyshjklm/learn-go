package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
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

func work(ch chan string) {
	// range + channel。channel是动态的。而range并不知道其长度。
	// 但根据当前状态进行取值，没有则阻塞。
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

	// 协程数量比数据多。一个channel，可以被多个协程竞争抢到
	// 就像一块骨头被几只狗抢。
	// 调度器会尽量保证调度的公平性。调度的算法及其实现
	// 这是生产-消费模式，也像用户通过请求到达网站入口，被调度到不同后端机器
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
