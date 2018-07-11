// just for test
package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type msgInfo struct {
	w http.ResponseWriter
	r *http.Request
	s string
}

var msgCh = make(chan msgInfo, 16)

func handleFunc(w http.ResponseWriter, r *http.Request) {
	msgCh <- msgInfo{w, r, string(time.Now().UnixNano())}
	log.Println("handler", r.URL.RequestURI())
	return
}

func processMsg() {
	for {
		select {
		case msg := <-msgCh:
			log.Printf("select %s\n", msg.r.URL.RequestURI())
			time.Sleep(2)
			_, err := fmt.Fprintf(msg.w, "hello %s\n", msg.r.URL.RequestURI())
			if err != nil {
				log.Println("response error")
			}
		default:
			time.Sleep(time.Millisecond * 10)
		}
	}
}

func main() {
	go processMsg()

	http.HandleFunc("/", handleFunc)
	http.ListenAndServe(":8080", nil)
}
