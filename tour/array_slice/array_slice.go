// mote types of slices.
package main 

import (
	"fmt"
)

func array_test() {
	var a [2]string
	a[0] = "hello"
	a[1] = "world"

	fmt.Println(a[0], a[1])
	fmt.Println(a)
}

func slice_test() {
	p := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("p ==", p)

	for i := 0; i < len(p); i++ {
		fmt.Printf("p[%d] == %d\n", i, p[i])
	}
	
	fmt.Println("p[1:4] ==", p[1:4])
	fmt.Println("p[:4] ==", p[:4])	// missing low index implies 0
	fmt.Println("p[4:] ==", p[4:])	// missing high index implies len(s)
}

func slice_make() {
	a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)
}

func printSlice(s string, x []int) {
	fmt.Printf("slice %s len=%d cap=%d %v\n", 
		s, len(x), cap(x), x)
}

func slice_nil() {
	var z []int
	fmt.Println(z, len(z), cap(z))

	if z == nil {
		fmt.Println("nil!")
	}
}

func slice_adding() {
	var a []int
	printSlice("a", a)

	// append works on nil slices.
	a = append(a, 0)
	printSlice("a", a)

	a = append(a, 1)
	printSlice("a", a)
	
	a = append(a, 2, 3, 4)
	printSlice("a", a)}

func main() {
	array_test()
	slice_test()
	slice_make()
	slice_nil()
	slice_adding()
}
