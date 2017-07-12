package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

type LineCounter int

// Version 1
func (line *LineCounter) Write(b []byte) (int, error) {
	// Count的第二个参数是分隔符
	*line = LineCounter(bytes.Count(b, []byte("\n")))
	return len(b), nil
}

// Version 2
// func (line *LineCounter) Write(b []byte) (int, error) {
// 	for _, v := range b {
// 		if v == '\n' {
// 			*line++
// 		}
// 	}
// 	return len(b), nil
// }

func main() {
	line := new(LineCounter)
	io.Copy(line, os.Stdin)
	fmt.Println(*line)
}

/*
go run LineCounter-3.go < LineCounter-3.go
40
wc -l LineCounter-3.go
      24 LineCounter-3.go
*/
