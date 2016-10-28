package main

import (
    "fmt"
    "unicode"
    "errors"
)

func nextInt(b []byte, i int) (int, int, error) {
    for ; i < len(b) && !unicode.IsDigit(rune(b[i])); i++ {
    }
    x := 0
    for ; i < len(b) && unicode.IsDigit(rune(b[i])); i++ {
        x = x*10 + int(b[i]) - '0'
    }
    if i < len(b) {
        return x, i, nil
    } else {
        return x, i, errors.New("outside offset")
    }
}

func main() {
    var x int
    var err error
    var src []byte = []byte("ab13579cd0ef8322d")
    for i := 0; i < len(src); {
        x, i, err = nextInt(src, i)
        if err == nil {
            fmt.Println(x)
        }
    }
    // Output:
    // 13579
    // 0
    // 8322
}
