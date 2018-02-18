package main

import (
	"fmt"
	"math"
)

// Point for 2-D point
type Point struct {
	X, Y float64
}

type Path []Point

// Distance compute distance of 2 point
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (path Path) Distance() float64 {
	var length float64
	for i := 1; i < len(path); i++ {
		length += path[i].Distance(path[i-1])
	}
	return length
}

func main() {
	var path Path
	path = []Point{{1, 2}, {3, 4}, {5, 6}}

	fmt.Println(path.Distance())
}
