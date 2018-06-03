package main

import "fmt"

type gopher struct {
	name string
	age  int
}

func main() {
	var m1 = make(map[string]*gopher)
	var m2 = make(map[string]*gopher)
	// make指定第二个参数时，这个参数长度与容量

	my := []gopher{
		{name: "tom", age: 11},
		{name: "jk", age: 12},
		{name: "alice", age: 13},
	}

	for i, r := range my {
		fmt.Printf("%+v, %+v\n", i, r)
		m1[r.name] = &r
		// r是内部的一个临时的变量，每次都获取得到的是一个相同的地址
		// 其值则是最后一次循环得到的用户信息
	}
	for k, v := range m1 {
		fmt.Println(k, v)
	}

	fmt.Println("\n--- m2 ---")
	for i, r := range my {
		fmt.Printf("%+v, %+v\n", i, r)
		m2[r.name] = &my[i]
	}

	for k, v := range m2 {
		fmt.Println(k, v)
	}
}
