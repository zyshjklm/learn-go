package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// Monitor struct for http counter
type Monitor struct {
	counter int64
}

// Run run the core logic
func (m *Monitor) Run() {
	for {
		time.Sleep(time.Second)
		m.counter++
	}
}

func (m *Monitor) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "counter:%d\n", m.counter)
}

func handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello Golang\n")
}

func main() {
	var m Monitor
	http.HandleFunc("/", handle)
	http.Handle("/monitor", &m)
	go m.Run()

	log.Fatal(http.ListenAndServe(":9090", nil))
}

/* usage:

go run http-monitor/main.go &

curl localhost:9090
hello Golang

curl localhost:9090/sdf
hello Golang

curl localhost:9090/monitor
counter:21
curl localhost:9090/monitor
counter:25

*/
