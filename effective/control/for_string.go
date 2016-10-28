package main

import "fmt"

func main() {
    for pos, char := range "China 中国國\x80语語" {
        // \x80 is an illegal UTF-8 encoding
        fmt.Printf("%U : starts at byte position %d.\n", char, pos)
    }
}

