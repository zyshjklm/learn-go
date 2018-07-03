package main

import (
	"fmt"
	"reflect"
)

type data struct {
	name     string
	password string
}

type Http struct {
	host  string
	agent string
	data
}

func main() {
	var h = Http{}
	t := reflect.TypeOf(h)
	fmt.Println(t)

	if name, ok := t.FieldByName("name"); ok {
		fmt.Printf("name:%+v\ntype:%+v\n\n", name, name.Type)
	}

	// 通过索引获取，Http的3个字段，Index分别是0,1,2
	// 故索引切片的第一个数字2表示的是data结构体。
	pwd := t.FieldByIndex([]int{2, 1})
	fmt.Printf("password:%+v\ntype:%+v\n\n", pwd, pwd.Type)
}
