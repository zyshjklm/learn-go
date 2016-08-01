// Ftoc print Fahrenheit-to-Celsius conversions
package main 

import (
	"fmt"
)

var p = f()

func f() *int {
	v := 1
	return &v
}

func main() {
	fmt.Printf("return: addr = %v, value = %d \n", p, *p)
	fmt.Printf("return: addr = %v, value = %d \n", p, *p)
	fmt.Println(f() == f())	
	fmt.Println(p == p)	
}

