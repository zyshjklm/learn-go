package main 

import (
	"fmt"
	"math"
)

// struct
type Vertex struct {
	X, Y float64
}

// method for struct
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// if define as (v Vertex), it has no effect when v is a Vertex
// when v is a value (non-pointer) type, the method sees
// a copy of the Vertex, and cannot mutate the original value.
func (v *Vertex) Scale( f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// built-in types
type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	v := &Vertex{3, 4}
	fmt.Println(v.Abs())
	v.Scale(5)
	fmt.Println(v.Abs())
	
	f := MyFloat(-math.Sqrt(2))
	fmt.Println(f, "\n", f.Abs())
}