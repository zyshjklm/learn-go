// practise of 3.1 - integer
package main

import (
    "fmt"
)

func main() {
    var x uint8 = 1<<1 | 1<<5
    var y uint8 = 1<<1 | 1<<2

    fmt.Printf("%08b\n", x)     // 00100010
    fmt.Printf("%08b\n", y)     // 00000110

    fmt.Printf("%08b\n", x&y)   // 00000010
    fmt.Printf("%08b\n", x|y)   // 00100110
    fmt.Printf("%08b\n", x^y)   // 00100100
    fmt.Printf("%08b\n", x&^y)  // 00100000

    for i := uint8(0); i < 8; i++ {
        if x&(1<<i) != 0 {
            fmt.Println(i)  // "1", "5"
        }
    }

    fmt.Printf("%08b\n", x<<1)  // 01000100
    fmt.Printf("%08b\n", x>>1)  // 00010001

    // use signed num for quantities.
    medals := []string{"gold", "silver", "bronze"}
    var length int = len(medals)

    for length--; length >= 0; length-- {
        fmt.Println(medals[length])
    }

    // dec, oct, hex
    o := 0666
    fmt.Printf("%d %[1]o %#[1]o\n", o)  // "438 66 0666"

    z := int64(0xdeadbeef)
    fmt.Printf("%d %[1]x %#[1]x\n", z)
    // "3735928559 deadbeef 0xdeadbeef"

    // rune literal
    ascii := 'a'
    unicode := '国'
    newline := '\n'

    fmt.Printf("%d %[1]c %[1]q\n", ascii)
    fmt.Printf("%d %[1]c %[1]q\n", unicode)
    fmt.Printf("%d %[1]q\n", newline)
}

