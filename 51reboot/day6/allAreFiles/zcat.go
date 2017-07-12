package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	r, err := gzip.NewReader(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("-- result --:")
	io.Copy(os.Stdout, r)
}
