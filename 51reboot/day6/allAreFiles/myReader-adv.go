package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"os"
	"archive/tar"
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
	// un tar
	tr := tar.NewReader(uncompress)
	hdr, err := tr.Next()
	info := hdr.FileInfo()

	// my reader
	myRd := NewMyReader(tr)

	// create a file fd as writer
	fd, err := os.Create(info.Name())
	if err != nil {
		log.Fatal("create file error!")
	}
	// multi writer
	mwr := io.MultiWriter(os.Stdout, fd)

	n, err := io.Copy(mwr, myRd)
	// io.Copy(os.Stdout, uncompress)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("copied num: %d; rd size: %d\n", n, myRd.cnt)
}
