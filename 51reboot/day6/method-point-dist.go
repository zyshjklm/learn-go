package main

import (
	"fmt"
	"math"
)

// Point for 2-D point
type Point struct {
	X, Y float64
}

func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// Distance compute distance of 2 point
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func main() {
	q := Point{1, 2}
	p := Point{4, 6}
	fmt.Println(Distance(p, q))
	fmt.Println(p.Distance(q)) // 5
}
