package main 

import (
	"fmt"
	"reflect"
)


type User struct {
	Id int
	Name string
	Age int
}

func (u User) Hello(name string) {
	fmt.Println("Hello", name, ", my name is", u.Name)
}


func Set(o interface{}) {
	v := reflect.ValueOf(o)

	if v.Kind() == reflect.Ptr && !v.Elem().CanSet() {
		fmt.Println("check ok in Set.")
		return 
	} else {
		v = v.Elem()
	}

	f := v.FieldByName("Name")
	if !f.IsValid() {
		fmt.Println("BAD")
		return
	}
	if f.Kind() == reflect.String {
		f.SetString("Alice")
	}
}




func main() {

	x := 123
	fmt.Println(x)

	v := reflect.ValueOf(&x)
	v.Elem().SetInt(999)
	fmt.Println(x)

	// User
	u := User{1, "Bob", 23}
	fmt.Println(u)

	Set(&u)
	fmt.Println(u)
	u.Hello("joe")

	v2 := reflect.ValueOf(u)
	mm := v2.MethodByName("Hello")

	args := []reflect.Value{ reflect.ValueOf("Kitty") }
	mm.Call(args)

}


