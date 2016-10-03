package main 

import (
	"fmt"
	"math"
)

type abser interface {
	Abs() float64
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	var a abser

	f := MyFloat(-math.Sqrt(2))
	v := Vertex{3, 4}

	a = f 	// a MyFloat implements abser
	fmt.Println(a.Abs())


	a = &v 	// a *Vertex implements abser
	fmt.Println(a.Abs())

	// v is a Vertex (not *Vertex). and does not implement abser
	//a = v
	}