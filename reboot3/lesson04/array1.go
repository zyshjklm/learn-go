package main

import "fmt"

func main() {
	var a [3]int
	fmt.Println(a[0])
	fmt.Println(a[len(a)-1])

	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}

	for i := range a {
		fmt.Printf("%d %d\n", i, a[i])
	}

	for _, v := range a {
		fmt.Printf("%d\n", v)
	}
	/*
		为什么数据的的index是从0开始的？
		p -> a
		p + 0
		p + 1 * sizeOf(type)
		p + 2 * sizeOf(type)

		string 16byte:
		8 : string address
		8 : len
	*/
}
