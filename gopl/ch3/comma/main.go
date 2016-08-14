// comma inserts commas in a non-negative decimal integer string.
package main

import (
    "fmt"
    "os"
    "bytes"
    "strings"
)

// run: ./comma 11234567
// output: 11,234,567

func comma0(s string) string {
    n := len(s)

    if n <= 3 {
        return s
    }
    return comma0(s[:n-3]) + "," + s[n-3:]
}


// exercise 3.10
func comma1(s string) string {
    return comma(s)
}

// exercise 3.11
func comma2(s string) string {
    var buf bytes.Buffer
    if s[0] == '-' {
        buf.WriteByte(s[0])
        s = s[1:]
    }

    if dot := strings.LastIndex(s, "."); dot >= 0 {
        buf.WriteString(comma(s[:dot])) // int
        buf.WriteString(s[dot:])        // .floatXXX
    } else {
        buf.WriteString(comma(s))
    }
    return buf.String()
}

func comma(s string) string {
    var buf bytes.Buffer
    
    length := len(s)
    loops := length/3
    // loop sub 1 when no remainder.
    if loops*3 == length {
        loops -= 1
    }
    head := length - loops * 3  
    idx_sta, idx_end := 0, head

    for i := loops; i >= 0; i-- {
        // fmt.Println("loops:", i, idx_sta, idx_end)
        if i != loops {
            buf.WriteByte(',')
        }
        buf.WriteString(s[idx_sta:idx_end])
        // fmt.Println(s[idx_sta:idx_end])

        idx_sta, idx_end = idx_end, idx_end + 3
    } 
    return buf.String()
}


func main() {
    for i := 1; i < len(os.Args); i++ {
        fmt.Printf(" %s\n", comma0(os.Args[i]))
        fmt.Printf(" %s\n", comma1(os.Args[i]))
        fmt.Printf(" %s\n", comma2(os.Args[i]))
    }
}

/*
 example:
./main -1234565.234  1234567 1234565.234
 -12,345,65.,234
 -12,345,65.,234
 -1,234,565.234
 1,234,567
 1,234,567
 1,234,567
 12,345,65.,234
 12,345,65.,234
 1,234,565.234
 */

