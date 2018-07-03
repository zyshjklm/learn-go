package main

import (
	"fmt"
	"reflect"
)

type Data struct {
	name     string
	password string
}

type Http struct {
	host  string
	agent string
	Data
}

func (h *Http) GetHost() string {
	return h.host
}

func (h *Http) GetAgent() string {
	return h.agent
}

func (d *Data) GetName() string {
	return d.name
}

func (d *Data) GetPass() string {
	return d.password
}

func main() {
	var h = Http{}
	t := reflect.TypeOf(&h)

	// range method
	fmt.Println("ptr:", t)
	for i := 0; i < t.NumMethod(); i++ {
		fmt.Printf("%v\n", t.Method(i))
	}

	fmt.Println("elem:", t.Elem())
	for i := 0; i < t.Elem().NumMethod(); i++ {
		fmt.Printf("%v\n", t.Elem().Method(i))
	}
}
