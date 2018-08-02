package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello golang...")
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":7878", nil))
}
