package main

import (
	"fmt"
	"log"
	"math"
	"os"
)

// IDistance interface
type IDistance interface {
	Distance() float64
}

// Point struct
type Point struct {
	X, Y float64
}

type Path []*Point

// Distance2Point compute distance of 2 point
func (p *Point) Distance2Point(q *Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// Distance compute distance to zero point
func (p *Point) Distance() float64 {
	return math.Hypot(p.X, p.Y)
}

// Distance distance of path
func (p Path) Distance() float64 {
	var sum float64
	for i := 0; i < len(p)-1; i++ {
		sum += p[i].Distance2Point(p[i+1])
	}
	return sum
}

// IDistance 接口为参数的函数。可以打印实现了该接口的类型
func print(p IDistance) {
	fmt.Println(p.Distance())
}

func main() {
	var path Path
	path = make([]*Point, 3)
	p1 := &Point{X: 1, Y: 2}
	p2 := &Point{X: 3, Y: 4}
	p3 := &Point{X: 5, Y: 6}
	path[0] = p1
	path[1] = p2
	path[2] = p3

	var i IDistance
	i = p1
	fmt.Println(i.Distance())
	i = p2
	fmt.Println(i.Distance())
	i = p3
	fmt.Println(i.Distance())

	print("path:", path)
	print("p p1:", p1)

	// Writer 接口
	f, err := os.Create("a.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	fmt.Fprintf(f, "hello\n")

}
