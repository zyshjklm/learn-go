package main

import (
    "fmt"
    "os"
    "strings"
    "time"
)

// use traditional for loop
func echo1() {
    var s, sep string
    fmt.Println("\nparameters:")

    for i := 1; i < len(os.Args); i++ {
        s += sep + os.Args[i]
        sep = " "
    }
    fmt.Println(s)
}

// use for range
func echo2() {
    s, sep := "", ""
    fmt.Println("\nparameters:")

    for _, arg := range os.Args[1:] {
        s += sep + arg
        sep = " "
    }
    fmt.Println(s)
}

// use strings.Join 
func echo3() {
    fmt.Println("\nparameters:")
    fmt.Println(strings.Join(os.Args[1:], " "))
}

func echoExercise1() {
    fmt.Println("\nfunc name and parameters:")
    fmt.Println(strings.Join(os.Args, " "))
}

func echoExercise2() {
    fmt.Println("\nparameters:\nindex\tparameter")
    for index, arg := range os.Args[1:] {
        fmt.Printf("%5d\t%s\n", index, arg)
    }
}

func main() {
    fmt.Println(os.Args)

    start := time.Now()
    echo1()
    fmt.Printf("%.8fs elapsed\n", time.Since(start).Seconds())

    echo2()

    start = time.Now()
    echo3()
    fmt.Printf("%.8fs elapsed\n", time.Since(start).Seconds())

    echoExercise1()
    echoExercise2()

}



