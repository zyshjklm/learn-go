package main

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
)

type Student struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
}

func (s Student) String() string {
	return fmt.Sprintf("name:%s,id:%d", s.Name, s.ID)
}

func marshal(x interface{}) string {
	t := reflect.TypeOf(x).Elem()
	v := reflect.ValueOf(x).Elem()
	buf := new(bytes.Buffer)

	fmt.Fprintf(buf, "{\n")
	for i := 0; i < t.NumField(); i++ {
		ft := t.Field(i)
		fv := v.Field(i)
		jsonkey := ft.Tag.Get("json")
		if jsonkey == "" {
			jsonkey = ft.Name
		}
		var jsonvalue string
		switch ft.Type.Kind() {
		case reflect.Int:
			jsonvalue = strconv.Itoa(int(fv.Int()))
		case reflect.String:
			jsonvalue = `"` + fv.String() + `"`
		}
		fmt.Fprintf(buf, "  \"%s\":%s,\n", jsonkey, jsonvalue)
	}
	fmt.Fprintf(buf, "}\n")
	return buf.String()
}

func main() {
	s := &Student{
		Name: "jungle",
		ID:   1,
	}
	fmt.Print(marshal(s))

	s1 := []string{"hello", "golang"}
	s2 := s1
	// fmt.Println(s1 == s2)
	// invalid operation: s1 == s2
	fmt.Println(reflect.DeepEqual(s1, s2)) //true
}
