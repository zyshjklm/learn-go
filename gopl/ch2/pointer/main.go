// Ftoc print Fahrenheit-to-Celsius conversions
package main 

import (
	"fmt"
)

var p = f()
var q = f()

func f() *int {
	v := 1
	return &v
}

func main() {
	fmt.Printf("return of p : addr = %v, value = %d \n", p, *p)
	fmt.Printf("return of q : addr = %v, value = %d \n", q, *q)
    fmt.Println("f() == f() :", f() == f())
    fmt.Println("p == q :", p == q)
}

