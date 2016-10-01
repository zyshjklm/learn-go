package main 

import (
	"fmt"
	"math"
)

// iterate 10 loops
func Sqrt1(x float64) float64 {
	z := 1.0	// assume z is the result.
	for i := 1; i < 10; i++ {
		z = z - (z*z - x) / (2*z)
		fmt.Printf("\ttemp: %4.8f\n", z)
	}
	return z
}

// iterate my delta
func Sqrt2(x float64) float64 {
	r := 1.0	// assume z is the result.
	for {
		newton := r - (r*r - x) / (2*r); 
		delta := math.Abs(newton - r) 
		if delta > 1e-10 {
			r = newton	
			fmt.Printf("\ttemp: %4.8f\n", r)
		} else {
			return r
		}
	}
	return r
}

func main() {
	for i := 1; i < 10; i++ {
		fmt.Printf("square root of %d is %4.8f, math result: %4.8f\n\n", 
			i, Sqrt1(float64(i)), math.Sqrt(float64(i)))
	}
	for i := 1; i < 10; i++ {
		fmt.Printf("square root of %d is %4.8f, math result: %4.8f\n\n", 
			i, Sqrt2(float64(i)), math.Sqrt(float64(i)))
	}
}