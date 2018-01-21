package main

import (
    "fmt"
)

var global string = "global"

func localFunc() {
    var local = "local"
    fmt.Println(local)
}

func main() {
    localFunc()
    fmt.Println(global)

    x := 1
    if true {
        var z int
        x, z = 100, 2
        // 一条赋值语句，多变量赋值，本句没有新变量；
        // x是main的变量，z是if的局部变量
        fmt.Println("x=", x, "z=", z)
        fmt.Println("&x=", &x)
    }
    fmt.Println(x)
    fmt.Println("&x=", &x)

    fmt.Println("--- style 2 ---")
    y := 1
    if true {
        y, z := 999, 3
        // y是声明的局部变量，if结束后就释放了
        fmt.Println("y=", y, "z=", z)
        fmt.Println("&y=", &y)
    }
    fmt.Println(y)
    fmt.Println("&y=", &y)

}
