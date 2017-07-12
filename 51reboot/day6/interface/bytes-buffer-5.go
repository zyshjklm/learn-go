package main

import (
	"bytes"
	"fmt"
	"io"
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
	l.Sum = bytes.Count(b, []byte("\n"))
	// for _, v := range b {
	// 	if v == '\n' {
	// 		l.Sum++
	// 	}
	// }
	return len(b), nil
}

func main() {
	b := new(ByteCounter)
	l := new(LineCounter)

	// A Buffer is a variable-sized buffer of bytes with Read and Write methods.
	buf := new(bytes.Buffer)
	buf.WriteString(`
hello gopher
1234
main new
`) // last ` in next line makes a new line.

	// from w to b and l
	w := io.MultiWriter(b, l)

	// from buf to w
	io.Copy(w, buf)
	fmt.Println("bytes:", b.Sum)
	fmt.Println("lines:", l.Sum)
}

/*
go run bytes-buf.go
28
4
*/
