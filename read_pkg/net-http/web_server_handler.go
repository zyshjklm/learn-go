package main

import (
    "log"
    "fmt"
    "net/http"
)

type String string

func (str String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "%s for URL.Path(%q)\n", str, r.URL.Path)
}

func main() {
    str := String("I'm jungle!")
    http.Handle("/string", str)
    log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
