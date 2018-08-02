package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func rootHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello golang...")
}

func helloHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "from hello...")
}

func worldHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "from world...")
}

func main() {
	s := http.Server{
		Addr:           ":7878",
		ReadTimeout:    20 * time.Second,
		WriteTimeout:   20 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/world", worldHandler)
	log.Fatal(s.ListenAndServe())
}
