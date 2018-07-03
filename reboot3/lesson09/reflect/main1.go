package main

import (
	"fmt"
	"reflect"
)

type X32 int32

func main() {
	var a X32 = 100
	t := reflect.TypeOf(a)
	fmt.Println(t.Name(), t.Kind())
}
