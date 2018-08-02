package main

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"runtime"
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

func logHTTP(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		p := runtime.FuncForPC(reflect.ValueOf(h).Pointer())
		name := p.Name()
		fmt.Println("handler func called = ", name)
		h(w, r)
	}
}

func main() {
	s := http.Server{
		Addr:           ":7878",
		ReadTimeout:    20 * time.Second,
		WriteTimeout:   20 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/hello", logHTTP(helloHandler))
	http.HandleFunc("/world", worldHandler)

	log.Fatal(s.ListenAndServe())
}
