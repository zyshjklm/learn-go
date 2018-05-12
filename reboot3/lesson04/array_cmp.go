package main

import "fmt"

func main() {
	a1 := [3]int{1, 2, 3}
	fmt.Println(a1)

	var a2 [3]int
	a2 = a1
	fmt.Println(a2)

	fmt.Println(a1 == a2)
	// true

	// var a3 = [2]int{1, 2}
	// fmt.Println(a1 == a3)
	// (mismatched types [3]int and [2]int)

	// 内存连续， 8byte
	for i := 0; i < 3; i++ {
		fmt.Printf("a1[%d]:%d, &a1[%d]:%p, &a1: %p\n",
			i, a1[i], i, &a1[i], &a1)
	}
	for i := 0; i < 3; i++ {
		fmt.Printf("a2[%d]:%d, &a2[%d]:%p, &a2: %p\n",
			i, a2[i], i, &a2[i], &a2)
	}
}

/*
a1[0]:1, &a1[0]:0xc420018180, &a1: 0xc420018180
a1[1]:2, &a1[1]:0xc420018188, &a1: 0xc420018180
a1[2]:3, &a1[2]:0xc420018190, &a1: 0xc420018180
a2[0]:1, &a2[0]:0xc4200181c0, &a2: 0xc4200181c0
a2[1]:2, &a2[1]:0xc4200181c8, &a2: 0xc4200181c0
a2[2]:3, &a2[2]:0xc4200181d0, &a2: 0xc4200181c0
地址间隔：Sizeof(int)
*/
