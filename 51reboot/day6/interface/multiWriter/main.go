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
	// 返回值是Writer接口，参数是多个Writer接口
	// 传递给w的值，会被复写到参数对应的多个Writer
	w := io.MultiWriter(b, l)

	// dest: w; src: os.Stdin
	io.Copy(w, os.Stdin)
	fmt.Println("bytes:", b.Sum)
	fmt.Println("lines:", l.Sum)
}

/*

go run multiWriter/main.go < LineCounter-3.go
bytes: 622
lines: 40

*/
