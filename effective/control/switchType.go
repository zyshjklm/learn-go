package main

import "fmt"

func getType(t interface{}) {
    switch t := t.(type) {
      // declaring a new var t with the same name of parameter
      // but a different type in each case.
    default:
        fmt.Printf("unexpected type %T\n", t)   // %T print whatever type t has
    case bool:
        fmt.Printf("boolean %t\n", t)
    case int:
        fmt.Printf("int %d\n", t)
    case *bool:
        fmt.Printf("pointer of boolean %t\n", *t)
    case *int:
        fmt.Printf("pointer of int %d\n", *t)
    }
}

func main() {
    var t interface{}
    t = 1
    getType(t)

    t = true
    getType(t)

    var x int = 3
    t = &x
    getType(t)

    var y bool = false
    t = &y
    getType(t)
}
