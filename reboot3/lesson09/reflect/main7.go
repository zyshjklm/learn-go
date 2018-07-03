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

func (h *Http) GetHost() string {
	return h.host
}

func (h *Http) GetAgent() string {
	return h.agent
}

func (h *Http) GetName() string {
	return h.data.name
}

func (h *Http) GetPass() string {
	return h.data.password
}

func main() {
	var h = Http{}
	t := reflect.TypeOf(&h)
	fmt.Println(t)

	// range methods
	for i := 0; i < t.NumMethod(); i++ {
		fmt.Printf("%+v\n", t.Method(i))
	}
}
