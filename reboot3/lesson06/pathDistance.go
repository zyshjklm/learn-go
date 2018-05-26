package main

import (
	"fmt"
	"math"
)

// Point for point
type Point struct {
	X, Y float64
}

type Path []Point

// Distance for p, q
func (p *Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func pathDistance(p Path) float64 {
	distance := 0.0
	if len(p) <= 1 {
		return 0.0
	}
	for i := len(p) - 1; i >= 1; i-- {
		distance += p[i].Distance(p[i-1])
	}
	return distance
}

func (p Path) Distance() float64 {
	distance := 0.0
	if len(p) <= 1 {
		return 0.0
	}
	for i := len(p) - 1; i >= 1; i-- {
		distance += p[i].Distance(p[i-1])
	}
	return distance
}

func main() {
	p := Path{Point{-1, 0}, {0, 0}, {3, 4}}

	fmt.Println(pathDistance(p))
	fmt.Println(p.Distance())
}
