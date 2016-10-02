package main 

import (
	"fmt"
)

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}
	v2 = Vertex{X: 1}	// y: 0 is implict
	v3 = Vertex{}		// x:0 and y:0
	p = &Vertex{5, 6}	// *Vertex type
)

func struct_test() {
	fmt.Println(Vertex{1, 2})

	v := Vertex{3, 4}
	fmt.Println(v)
	v.X = 7
	fmt.Println(v)

	// pointer to struct
	v = Vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)
}


func main() {
	struct_test()
	fmt.Println(v1, v2, v3)
	fmt.Println(p)
}

