package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

type XorWriter struct {
	w io.Writer
	x byte
}

// Write writes len(p) bytes from p to the underlying data stream.
func (x *XorWriter) Write(p []byte) (int, error) {
	p1 := make([]byte, len(p))
	copy(p1, p) // copy(dst, src)
	for i, b := range p1 {
		p1[i] = b ^ x.x
	}
	// from p -> p1(compute p1) -> x.w.Write()
	return x.w.Write(p1)
}

func NewXorWriter(w io.Writer, x byte) *XorWriter {
	return &XorWriter{
		w: w,
		x: x,
	}
}

type XorReader struct {
	r io.Reader
	x byte
}

// Read reads up to len(p) bytes into p
func (x *XorReader) Read(p []byte) (int, error) {
	// return x.r.Read(p)
	buf := make([]byte, len(p))
	n, err := x.r.Read(buf)
	fmt.Printf("read: ")
	for i := 0; i < n; i++ {
		fmt.Printf("%v ", buf[i])
	}
	fmt.Println()
	copy(p, buf) // copy(dst, src)
	return n, err
}

func NewXorReader(r io.Reader, x byte) *XorReader {
	return &XorReader{
		r: r,
		x: x,
	}
}

func main() {
	buf := new(bytes.Buffer)
	ori := "hello golang aabbccdd"
	fmt.Println("str:", ori)
	x := NewXorWriter(buf, 'a')
	fmt.Println("buf:", buf.Bytes())

	io.WriteString(x, ori)
	fmt.Println("xor wr:", buf.Bytes())

	x1 := NewXorReader(buf, 'a')
	fmt.Println("xor rd:", buf.Bytes())
	fmt.Println("-- before copy()--")
	n, _ := io.Copy(os.Stdout, x1)
	fmt.Println("end:", n)

}
