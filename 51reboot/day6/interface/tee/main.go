package main

import (
	"fmt"
	"io"
	"os"
)

type LineCounter struct {
	Sum int
}
type ByteCounter struct {
	Sum int
}

func (b *ByteCounter) Write(p []byte) (int, error) {
	b.Sum += len(p)
	return len(p), nil
}

func (l *LineCounter) Write(b []byte) (int, error) {
	for _, v := range b {
		if v == '\n' {
			l.Sum++
		}
	}
	return len(b), nil
}

func main() {
	b := new(ByteCounter)
	l := new(LineCounter)
	// read from os.Stdin, write to b; and return a reader
	teeRd := io.TeeReader(os.Stdin, b)

	// dst: l; src: teeRd
	io.Copy(l, teeRd)
	fmt.Println(b.Sum)
	fmt.Println(l.Sum)
}

/*
go run tee/main.go < LineCounter-3.go
622
40
*/
