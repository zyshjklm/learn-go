package main

import "fmt"

// Point for 2-D point
type Point struct {
	X, Y float64
}

// Distance compute distance of 2 point
func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

func main() {
	//  direct pointer
	p0 := &Point{1, 2}
	fmt.Println(p0)
	p0.ScaleBy(2)
	fmt.Println(p0)

	// 声明结构体后，再用指针指向
	p1 := Point{1, 2}
	p2 := &p1
	fmt.Println("", p1)
	fmt.Println(p2)
	p2.ScaleBy(3)
	fmt.Println(p2)

	// 使用结构体，隐式取地址
	p3 := Point{1, 2}
	fmt.Println(p3)
	p3.ScaleBy(4) // 等价于 (&p3).ScaleBy(2)
	fmt.Println(p3)

	// 总结：定义时使用指针，使用时使用对象，隐式调用。
}
