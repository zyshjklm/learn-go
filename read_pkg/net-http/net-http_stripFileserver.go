package main

import (
    "net/http"
    "log"
)

func main() {
    // to server a directory on disk(/tmp/) under an alternate URL
    // path(/tmpfiles/), use StripPrefix to modify the request RUL's
    // path before the FileServer sees it:
    http.Handle("/home/", http.StripPrefix("/home/",
        http.FileServer(http.Dir("/Users/user/"))))

    log.Fatal(http.ListenAndServe(":8080", nil))
}

