package main

import (
	"fmt"
	"io"
	"os"
)

type ByteCounter struct {
	Sum int
}

func (b *ByteCounter) Write(p []byte) (int, error) {
	b.Sum += len(p)
	return b.Sum, nil
}

func main() {
	b := new(ByteCounter)
	io.Copy(b, os.Stdin)
	fmt.Println(b.Sum)
}

/*
go run ByteCounter-struct-2.go << EOF
heredoc> hello
heredoc> EOF
6

echo     hello | go run ByteCounter-struct-2.go
6
echo -n  hello | go run ByteCounter-struct-2.go
5
go run ByteCounter-struct-2.go < ByteCounter-struct-2.go
493
*/
