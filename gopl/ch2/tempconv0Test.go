// test tempconv
package main

import (
    "github.com/jungle85gopy/learn-go/gopl/ch2/tempconv0"
    "fmt"
)

func main() {
    fmt.Printf("%g\n", tempconv.BoilingC - tempconv.FreezingC)
    boilingF := tempconv.CToF(tempconv.BoilingC)
    fmt.Println(tempconv.BoilingC.String())
    fmt.Printf("%g\n", boilingF - tempconv.CToF(tempconv.FreezingC))


}

