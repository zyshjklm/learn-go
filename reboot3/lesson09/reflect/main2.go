package main

import (
	"fmt"
	"reflect"
)

type X32 int32

func main() {
	var a X32 = 100
	t := reflect.TypeOf(a)
	fmt.Printf("var:\n\tName: %+v, Kind: %+v\n", t.Name(), t.Kind())

	tp := reflect.TypeOf(&a)
	fmt.Printf("pointer:\n\tKind: %+v, Elem: %+v\n", tp.Kind(), tp.Elem())
}
