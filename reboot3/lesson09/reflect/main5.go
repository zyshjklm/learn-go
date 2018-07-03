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
	t := reflect.TypeOf(&h)
	fmt.Println(t)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	// 遍历字段
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		printFiled("", f)
		if f.Anonymous {
			// anonymous embedded struct
			for x := 0; x < f.Type.NumField(); x++ {
				sf := f.Type.Field(x)
				printFiled("-- ", sf)
			}
		}
	}
}

func printFiled(prefix string, sf reflect.StructField) {
	fmt.Printf("%stype:%+v,name:%+v,index:%+v,offset:%+v\n",
		prefix, sf.Type, sf.Name, sf.Index, sf.Offset)
}
