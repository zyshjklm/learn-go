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

func logHandler(h http.Handler) http.Handler {
	temp := func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Handler called - %T\n", h)
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(temp)
}

func main() {
	handle := myhandle{}
	handleHello := myhello{}
	handleWorld := myhello{}

	s := http.Server{
		Addr:           ":7878",
		ReadTimeout:    20 * time.Second,
		WriteTimeout:   20 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	http.Handle("/", &handle)
	http.Handle("/hello", logHandler(&handleHello))
	http.Handle("/world", &handleWorld)

	log.Fatal(s.ListenAndServe())
}
