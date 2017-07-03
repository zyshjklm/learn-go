package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func v1() {
	buf := make([]byte, 1024)
	for {
		n, err := os.Stdin.Read(buf)
		if err != nil {
			return
		}
		os.Stdout.Write(buf[:n])
	}
}

func v2() {
	var f *os.File
	var err error

	if len(os.Args) > 1 {
		f, err = os.Open(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
	} else {
		f = os.Stdin
	}
	io.Copy(os.Stdout, f)
}
func main() {
	// v1()
	fmt.Println("exit from v1(), star v2()...")
	v2()
}
