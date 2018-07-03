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

	t0 := reflect.TypeOf(&h)
	fmt.Println(t0)

	if t0.Kind() == reflect.Ptr {
		t0 = t0.Elem()
	}
	fmt.Println(t0)

	t1 := reflect.TypeOf(h)
	fmt.Println(t1)
}
