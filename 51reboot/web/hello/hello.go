package main

import (
	"io"
	"log"
	"net/http"
)

// Hello usage:
// curl localhost:8090/hello
func Hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello http\n")
}

func main() {
	http.HandleFunc("/hello", Hello)
	log.Fatal(http.ListenAndServe(":8090", nil))
}
