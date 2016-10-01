// (Package doc comment intentionally malformed to demonstrate golint.)
package main

import "fmt"
import "github.com/jungle85gopy/learn-go/gopl/ch2/popcount"

func main() {
    ori := [...]int{1,3,7,15,31,63,127, 255, 511,1023,2047,4095,8091,65535}
    for _, i := range ori {
        fmt.Printf("%d: %d\n", i, popcount.PopCount(uint64(i)))
    }
}
