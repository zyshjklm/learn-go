package main

import (
	"fmt"
	"math"
)

// Point for 2-D point
type Point struct {
	X, Y float64
}

// Distance compute distance of 2 point
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func Distance(path []Point) (length float64) {
	for i := 1; i < len(path); i++ {
		length += math.Hypot(path[i-1].X-path[i].X, path[i-1].Y-path[i].Y)
	}
	return
}

func Distance1(path []Point) (length float64) {
	for i := 1; i < len(path); i++ {
		length += path[i].Distance(path[i-1])
	}
	return
}

func main() {
	path := []Point{{1, 2}, {3, 4}, {5, 6}}

	fmt.Println(Distance(path))
	fmt.Println(Distance1(path))
}
