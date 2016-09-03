package main 

import (
	"fmt"
)

func appendInt(sliceX []int, y ...int) []int {
	var result []int
	rlen := len(sliceX) + len(y)

	if rlen <= cap(sliceX) {
		// there is room to grow. extend the slice.
		result = sliceX[:rlen]
	} else {
		// there is insuffient space. Allocate a new array.
		// grow by doubling. 
		rcap := rlen
		if rcap < 2 * len(sliceX) {
			rcap = 2 * len(sliceX)
		}
		result = make([]int, rlen, rcap)
		copy(result, sliceX)
	}

	// result[len(sliceX)] = y
	copy(result[len(sliceX):], y)
	return result
}

func main() {
	var x, y []int

	for i := 0; i < 10; i++ {
		y = appendInt(x, i, i+1, i+2)
		fmt.Printf("%d cap=%3d\t%v\n", i, cap(y), y)
		x = y
	}
}