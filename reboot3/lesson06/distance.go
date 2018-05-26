package main

import (
	"fmt"
	"math"
)

// Point for point
type Point struct {
	X, Y float64
}

// Distance for p, q
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// Distance for p, q
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func main() {
	p := Point{0, 0}
	q := Point{3, 4}

	fmt.Println(Distance(p, q)) // function call

	fmt.Println(p.Distance(q)) // method call
}
