package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"os"
)

type MyReader struct {
	reader io.Reader
	cnt    int
}

func NewMyReader(rd io.Reader) *MyReader {
	return &MyReader{
		reader: rd,
		cnt:    0,
	}
}

func (rd *MyReader) Read(b []byte) (int, error) {
	n, err := rd.reader.Read(b)
	rd.cnt += n
	return n, err
}

// only process a tar.gz with one file in current path. 
func main() {
	// un gzip
	uncompress, err := gzip.NewReader(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	myRd := NewMyReader(uncompress)
	n, err := io.Copy(os.Stdout, myRd)
	// io.Copy(os.Stdout, uncompress)
	fmt.Printf("copied num: %d; rd size: %d\n", n, myRd.cnt)
}
