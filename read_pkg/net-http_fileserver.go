package main

import (
    "net/http"
    "log"
)

func main() {
    // simple static webserver
    log.Fatal(http.ListenAndServe(":8080",
        http.FileServer(http.Dir("/usr/share/doc"))))
}

