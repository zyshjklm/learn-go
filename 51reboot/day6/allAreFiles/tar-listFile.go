package main

import (
	"archive/tar"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	tr := tar.NewReader(os.Stdin)
	for {
		hdr, err := tr.Next()
		if err != nil {
			return
		}
		fmt.Println(hdr.Name)
		io.Copy(ioutil.Discard, tr)
	}
}
