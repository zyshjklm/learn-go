package main

import "fmt"

func main() {
	var flist []func()
	var i int

	fmt.Println("\nbefore set flist, i=", i, &i)
	for i = 0; i < 3; i++ {
		flist = append(flist, func() {
			fmt.Println(i, &i)
		})
	}
	fmt.Println(" after set flist, i=", i, &i)
	// flist中的3个函数，所使用的i，都是一个地址
	for _, f := range flist {
		f()
	}
	i = 0
	var flist2 []func()
	fmt.Println("\nbefore set flist, i=", i, &i)
	for i = 0; i < 3; i++ {
		j := i // break the hole
		flist2 = append(flist2, func() {
			fmt.Println(j, &j)
		})
	}
	fmt.Println(" after set flist, i=", i, &i)
	for _, f := range flist2 {
		f()
	}
}
