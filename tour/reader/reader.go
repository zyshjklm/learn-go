package main 

import (
	"fmt"
	"io"
	"strings"
)

func string_reader() {
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)
	for {
		// Read() read(consume) from r,  output 8 bytes at a time to b. 
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}

func main() {
	string_reader()
}

