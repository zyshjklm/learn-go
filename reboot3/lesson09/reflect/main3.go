package main

import (
	"fmt"
	"reflect"
)

type X32 int32

func main() {
	var a X32 = 100

	v := reflect.ValueOf(a)
	fmt.Printf("type: %+v, value:%+v\n", v.Type(), v)
}
