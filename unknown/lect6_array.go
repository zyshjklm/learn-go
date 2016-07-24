package main 

import "fmt"

func main() {
	var a = [2]int{}
	var b = [3]int{1}	// [1, 0, 0]
	// [3]int 整体做为一个类型

	// error of : b = a
	fmt.Println(a, b)
	c := [20]int{18:1}	// offset 18为1，即倒数第2个数
	fmt.Println(c)

	d := [...]int{1,2,3,4,5}	// ... 由后面的长度决定
	fmt.Println(d)

	e := [...]int{5:9}
	fmt.Println(e)	// [0 0 0 0 0 9]


	// 指向数组的指针. *[6]int means pointer for [6]int
	var p *[6]int = &e
	// 此处，定义p和e时的数组长度值须相同。
	fmt.Println(p)	// &[0 0 0 0 0 9]

	// 指针数组
	x, y := 1, 2
	arr := [...]*int{&x, &y}
	fmt.Println(arr)

	// 数组间可以使用==, !=进行比较。但不能进行>,<比较
	m := [2]int{1,2}
	n := [2]int{1,2}
	fmt.Println(m == n, "\n")	// true

	// new
	q := new( [10]int )
	q[2] = 2
	fmt.Println(q)

	r := [10]int{}
	r[2] = 3
	fmt.Println(r)

	// 2-D
	op1 := [2][3]int {
		{1,1,1},
		{2,2,2}}
	op2 := [2][3]int {
		{1:1},
		{2:2}}
	fmt.Println(op1)
	fmt.Println(op2)

	for i := 0; i < 3; i++ {
		vv := 2
		fmt.Println(&vv)
	}
}



