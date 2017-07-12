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
