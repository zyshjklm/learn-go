package main

import (
	"fmt"
	"io"
	"os"
)

type ByteCounter int

func (b *ByteCounter) Write(p []byte) (int, error) {
	// 将入参的长度值转换为ByteCounter类型。
	*b += ByteCounter(len(p))
	return len(p), nil
}

func main() {
	b := new(ByteCounter)
	io.Copy(b, os.Stdin)
	fmt.Println(*b)
}

/*
 go run ByteCounter-interface-4.go <<EOF
heredoc> golang
heredoc> EOF
7

echo     hello | go run ByteCounter-interface-4.go
6
echo -n  hello | go run ByteCounter-interface-4.go
5

go run ByteCounter-interface-4.go < ByteCounter-interface-4.go
557
*/
