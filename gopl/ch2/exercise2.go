// exercise 2.2
package main

import (
    "os"
    "fmt"
    "bufio"
    "strconv"

    "github.com/jungle85gopy/learn-go/gopl/ch2/tempconv"
)

func main() {
    args := os.Args[1:]

    if len(args) < 1 {
        input := bufio.NewScanner(os.Stdin)
        for input.Scan() {
            f := str2f(input.Text())
            transfer(f)
        }
    } else {
        for _, arg := range args {
            f := str2f(arg)
            transfer(f)
        }
    }
}

func str2f(s string) float64 {
    f, err := strconv.ParseFloat(s, 64)
    if err != nil {
        fmt.Fprintf(os.Stderr, "str2f: %v\n", err)
    }
    return f
}

func transfer(f float64) {
    fmt.Printf("str2f :%f\n", f)
    fmt.Printf(" = %f\n", tempconv.CToF(tempconv.Celsius(f)))
}

