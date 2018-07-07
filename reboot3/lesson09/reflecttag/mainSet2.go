package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type Gopher struct {
	Name string
	age  string
}

func main() {
	var g = Gopher{}
	v := reflect.ValueOf(&g).Elem()

	name := v.FieldByName("Name")
	age := v.FieldByName("age")

	fmt.Printf("type:%+v, canAddr:%+v, canSet:%+v\n",
		name.Type(), name.CanAddr(), name.CanSet())

	fmt.Printf("type:%+v, canAddr:%+v, canSet:%+v\n",
		age.Type(), age.CanAddr(), age.CanSet())

	g.Name = "cc"
	*(*string)(unsafe.Pointer(age.UnsafeAddr())) = "20"

	fmt.Printf("name:%+v, age%+v\n", name, age)
	fmt.Printf("%+v\n", g)
}
