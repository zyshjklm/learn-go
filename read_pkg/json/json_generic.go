package main

import "fmt"
import "math"


func check_type(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("twice i is:", v*2)
	case float64:
		fmt.Println("the reciprocal of i is:", 1/v)
	case string:
		h := len(v) / 2
		fmt.Println("i swapped by halves is:", v[h:] + v[:h])
	default:
		fmt.Println("unknown types of i!")
	}
}

func main() {
    var i interface{}

    i = "a string"
    check_type(i)

    i = 2011
    check_type(i)

    i = 2.718
    check_type(i)

    r := i.(float64)
    fmt.Println("the circle's area:", math.Pi*r*r)

}

