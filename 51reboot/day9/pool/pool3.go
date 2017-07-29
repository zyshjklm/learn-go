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

// 1 只要不被close，可以永远发送数据和接受数据
// 2 如果channel中没有数据，接收方会阻塞
// 3 如果没有人正在等等channel的数据，发送方会阻塞。
//	   因此创建好channel后不能马上向其写数据，而是要先用协程并使用channel
// 	   如 go work(ch)，这样才会有消费方
// 4 close掉channel时，所有使用它的协程在取数据时得到默认值。
func work(ch chan string) {
	// channel的长度是动态的。不能当成是数组。
	// 本次循环可能没有数据，则会导致channel阻塞
	// 直到收到有数据，则唤醒本次循环。
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
