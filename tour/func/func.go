package main 

import (
	"fmt"
	"math"
)


func func_example() {
	// functions are values too.
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(3, 4))
}

// closure
// return a func defined as: func(x int) int {}
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func test_adder() {
	pos, neg := adder(), adder()
	for i := 0; i < 5; i++ {
		fmt.Println(pos(i), neg(-2*i))
	}
}

/*
output:
0 0
1 -2
3 -6
6 -12
10 -20
*/

func main() {
	func_example()
	test_adder()
}
