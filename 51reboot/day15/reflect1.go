package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
}

func (s Student) String() string {
	return fmt.Sprintf("name:%s,id:%d", s.Name, s.ID)
}

func print(v interface{}) {
	// 参数可以是任意类型。因此函数在运行时需要动态的获取v的类型
	t := reflect.TypeOf(v)
	fmt.Println("kind:", t.Kind())    // Kind()获取数据类型，此处是: ptr
	t = t.Elem()                      // 获取具体元素。
	fmt.Println("name:", t.Name())    // 获取名字
	fmt.Println("path:", t.PkgPath()) // 获取全路径

	fmt.Println("\n--- iter of field:")
	for i := 0; i < t.NumField(); i++ {
		filed := t.Field(i)
		fmt.Println("\t", filed)
		fmt.Println("\tjson-key:", filed.Tag.Get("json"))
	}
	// 根据具体的字段名来获取信息。
	f, _ := t.FieldByName("Name")
	fmt.Printf("\nFieldByName:%#v\n\n", f)

	for i := 0; i < t.NumMethod(); i++ {
		method := t.Method(i)
		fmt.Println(method.Name) // 值为: String
	}
	// value
	vl := reflect.ValueOf(v).Elem()
	vlfiled := vl.FieldByName("Name")
	fmt.Println(vlfiled.String())

	// 获取方法名，调用方法
	method := vl.MethodByName("String")
	ret := method.Call(nil)
	fmt.Println(ret[0].String())
}

func main() {
	// 反射，动态的操纵类型，类型，字段等
	s := &Student{
		Name: "jungle",
		ID:   1,
	}
	print(s)
}
