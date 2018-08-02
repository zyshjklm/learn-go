package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "hello golang...")
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServeTLS(":7878", "cert.pem", "key.pem", nil))
	// log.Fatal(http.ListenAndServe(":7878", nil))
}
