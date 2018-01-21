package main

import (
    "fmt"
)

func main() {
    var x int
    var p *int
    var pp **int
    //p = 0x0123
    // ./main.go:10: cannot use 291 (type int) as type *int in assignment
    p = &x
    pp = &p

    fmt.Println(" p = ", p)
    fmt.Println("&x = ", &x)
    fmt.Println("pp = ", pp)
}
