package main

import (
	"fmt"
	"reflect"
)

type X32 int32

func main() {
	var a X32 = 100

	va := reflect.ValueOf(&a)
	fmt.Printf("type:%+v, canAddr:%+v, canSet:%+v\n",
		va.Type(), va.CanAddr(), va.CanSet())

	vp := reflect.ValueOf(&a).Elem()
	fmt.Printf("type:%+v, canAddr:%+v, canSet:%+v\n",
		vp.Type(), vp.CanAddr(), vp.CanSet())

	var b X32 = 50
	vb := reflect.ValueOf(b)

	fmt.Println(vp)
	vp.Set(vb)
	fmt.Println(vp)

	//vp.Set(X32(99))
	//cannot use X32(99) (type X32) as type reflect.Value in argument to vp.Set
	//fmt.Println(vp)
}
