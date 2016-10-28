package main

import "fmt"

func main() {
    // array
    a := [...]string   {"no error", "Eio", "invalid argument"}
    // slice
    s := []string   {"no error", "Eio", "invalid argument"}
    m := map[int]string{1: "no error", 2: "Eio", 3: "invalid argument"}

    fmt.Println(a)
    fmt.Println(s)
    fmt.Println(m)
}
