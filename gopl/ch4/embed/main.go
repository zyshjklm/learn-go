package main 

import (
	"fmt"
)

type Point struct {
	X, Y int
}

type Circle struct {
	Point
	Radius int
}

type  Wheel struct {
	Circle
	Spokes int
}

func main() {
	w1 := Wheel{
		Circle: Circle{
			Point: Point{X: 18, Y: 18},
			Radius: 15,
		},
		Spokes: 20,
	}

	fmt.Printf("%#v\n", w1)
	// output:
	w1.X = 19
	fmt.Printf("%#v\n", w1)

	w2 := Wheel{Circle{Point{28, 28}, 25,}, 21}
	fmt.Printf("%#v\n", w2)
}