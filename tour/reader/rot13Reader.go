// rot13 reader
package main 

import (
	"io"
	"os"
	"strings"
	"fmt"
)

type rot13Reader struct {
	r io.Reader
}

// ROT13 refer :
// 	https://en.wikipedia.org/wiki/ROT13
func (rot rot13Reader) Read(b []byte) (int, error) {
	n, err := rot.r.Read(b)
	if n <= 0 {
		return n, err
	}

	fmt.Println("len:", len(b), "cap:", cap(b))
	// if you use for i, c := range b {}
	// the loop will be 32768 times.

	for i := 0; i < n; i++ {
	//	fmt.Printf("ori:%v", b[i])
		b[i] = rot.cryptRot13(b[i])
	//	fmt.Println(" new:%v", b[i])
	}
	return n, err
}

func (rot13Reader) cryptRot13(c byte) byte {
	if c <= 'Z' && c >= 'A' {
		c = (c - 'A' + 13) % 26 + 'A'
	} else if c <= 'z' && c >= 'a' {
		c = (c - 'a' + 13) % 26 + 'a'
	}
	return c
}


func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!\n")
	r := rot13Reader{s}

	// output: 'You cracked the code!'
	io.Copy(os.Stdout, &r)
}

