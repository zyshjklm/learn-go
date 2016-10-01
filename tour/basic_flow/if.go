package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow1(x, n, limit float64) float64 {
	// v is only in scope if the if statement.
	if v := math.Pow(x, n); v < limit {
		return v
	}
	// if using varible v here. compiler return: undefined: v
	return limit
}

func pow2(x, n, limit float64) float64 {
	if v := math.Pow(x, n); v < limit {
		return v
	} else {
		// v is available in else here.
		fmt.Printf("%g >= %g\n", v, limit)
	}
	// can't use v here.
	return limit
}

func main() {	
	fmt.Println(sqrt(2), sqrt(-4))
	fmt.Println(pow1(3, 2, 10), pow1(3, 3, 20))
	fmt.Println(pow2(3, 2, 10), pow2(3, 3, 20))

}