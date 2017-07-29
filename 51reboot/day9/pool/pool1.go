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
	url := <-ch
	printURLStatus(url)
}

func main() {
	var ch = make(chan string)
	go work(ch)

	ch <- "http://www.baidu.com"
	time.Sleep(time.Second)
	close(ch)
}
