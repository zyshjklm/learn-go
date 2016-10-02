package main 

import (
	"fmt"
)

func main() {
	i, j := 42, 2701

	p := &i 			// pointer to i
	fmt.Println(*p)		// read i through pointer
	*p = 32				// set  i through pointer
	fmt.Println(i)		// see the new value of i

	p = &j			// pointer to j
	*p = *p / 37	// divide j through the pointer
	fmt.Println(j)	// set the new value of j
}