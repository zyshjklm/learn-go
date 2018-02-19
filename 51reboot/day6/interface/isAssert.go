package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	var w io.Writer
	w = os.Stdout
	f := w.(*os.File)
	fmt.Println(f)
	_ = f

	wa, ok := w.(*bytes.Buffer)
	if ok {
		fmt.Println("assert ok")
	} else {
		fmt.Println("assert err")
	}
	_ = wa
	// use x, ok for assert

	c := w.(*bytes.Buffer)
	fmt.Println(c)
	_ = c

	/*
	   panic: interface conversion: io.Writer is *os.File, not *bytes.Buffer

	   goroutine 1 [running]:
	   main.main()
	   	iAssert.go:17 +0x7c
	*/
}
