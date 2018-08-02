package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type myhandle struct{}
type myhello struct{}
type myworld struct{}

func (h *myhandle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello root")
}

func (h *myhello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello from golang")
}

func (h *myworld) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}

func main() {
	handle := myhandle{}
	handleHello := myhello{}
	handleWorld := myworld{}
	s := http.Server{
		Addr:           ":7878",
		ReadTimeout:    20 * time.Second,
		WriteTimeout:   20 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	http.Handle("/", &handle)
	http.Handle("/hello", &handleHello)
	http.Handle("/world", &handleWorld)
	log.Fatal(s.ListenAndServe())
}
