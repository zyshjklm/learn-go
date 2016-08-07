// practise of 3.2 - float
package main

import (
    "fmt"
    "math"
)

const (
    e = 2.71828
    Avogadro = 6.02214129e23
    Planck = 6.62606957e-34
)

func main() {
    var f float32 = 1 << 24
    fmt.Println(f == f+1)   // true !!!
    // float32 is not large.

    fmt.Println(e, Avogadro, Planck)
    fmt.Printf("%f, %f, %f\n", e, Avogadro, Planck)

    // limits
    fmt.Println("\nmax of float32 :", math.MaxFloat32,
        "\nmax of float64 :", math.MaxFloat64)

    // format for float
    for x := 0; x < 8; x++ {
        fmt.Printf("x = %d e^x = %8.3f\n", x, math.Exp(float64(x)))
    }

    // about zero
    var z float64
    fmt.Println(z, -z, 1/z, -1/z, z/z)
    // 0 -0 +Inf -Inf NaN

    // NaN
    nan := math.NaN()
    fmt.Println(nan, math.IsNaN(nan))   // NaN true
    fmt.Println(nan == nan, nan < nan, nan > nan)   // false false false
      

}

