// comma inserts commas in a non-negative decimal integer string.
package main

import (
    "fmt"
    "os"
)


func comma0(s string) string {
    n := len(s)

    if n < 3 {
        return s
    }
    return comma(s[:n-3]) + "," + s[n-3:]
}

func main() {
    for i := 1; i < len(os.Args); i++ {
        fmt.Printf(" %s\n", comma0(os.Args[i]))
    }
}
