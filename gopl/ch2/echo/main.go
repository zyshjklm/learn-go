// echo prints its command-line arguments.
package main

import (
    "fmt"
    "flag"
    "strings"
)

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

func main() {
    flag.Parse()
    fmt.Print(strings.Join(flag.Args(), *sep))
    //fmt.Printf("%v\n", *n)
    if !*n {
        fmt.Println()
    }
}

