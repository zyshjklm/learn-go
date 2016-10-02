// mote types of range.
package main 

import (
	"fmt"
)

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}


func main() {
	// index and value
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	pow1 := make( []int, 10)
	// index only
	for i := range pow1 {
		pow1[i] = 1 << uint(i)
	}
	
	// value only
	for _, v := range pow1 {
		fmt.Printf("%d\n", v)
	}

}
