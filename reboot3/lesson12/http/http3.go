package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type myhandle struct{}

func (h *myhandle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello golang from myhandle...")
}

func main() {
	handle := myhandle{}

	s := http.Server{
		Addr:           ":7878",
		ReadTimeout:    20 * time.Second,
		WriteTimeout:   20 * time.Second,
		MaxHeaderBytes: 1 << 20,
		Handler:        &handle,
	}
	log.Fatal(s.ListenAndServe())
}
