package main

import (
	"fmt"
	"io"
	"os"
)

type ByteCounter int

func (b *ByteCounter) Write(p []byte) (int, error) {
	*b += ByteCounter(len(p))
	return len(p), nil
}

func main() {
	b := new(ByteCounter)

	io.Copy(b, os.Stdin)
	fmt.Println(*b)
}

/*
echo "hello" | go run byteCounter.go
6

go run byteCounter.go < byteCounter.go
249
*/
