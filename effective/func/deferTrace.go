package main

import "fmt"

func trace(s string) string {
    fmt.Printf("\n---- trace: entering: %s\n", s)
    return s
}

func untrace(s string) {
    fmt.Printf("-- untrace: leaving : %s\n", s)
}

func funcA() {
    var funcName = "funcA"
    defer untrace(trace(funcName))
    // the arguments to deferred func are evaluated when the defer executes.
    // so the tracing routine can set up the argument to the untracing routine.

    // do sth here...
    fmt.Println("working in ", funcName)
}

func funcB() {
    var funcName = "funcB"
    defer untrace(trace(funcName))

    // do sth here...
    fmt.Println("working in ", funcName)
}


func funcC() {
    var funcName = "funcC"
    defer untrace(trace(funcName))

    // do sth here...
    fmt.Println("working in ", funcName)

    funcB()
}


func main() {
    funcA()
    funcB()
    funcC()
}
