package main

import (
	"archive/tar"
	"io"
	"log"
	"os"
	"path/filepath"
)

func unTar(base string, r io.Reader) error {
	// create a reader for tar
	tr := tar.NewReader(os.Stdin)
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return nil
		}
		fullPath := filepath.Join(base, hdr.Name)
		info := hdr.FileInfo()
		// for dir
		if info.IsDir() {
			os.MkdirAll(fullPath, 0755)
			continue
		}
		// for file
		dir := filepath.Dir(fullPath)
		os.MkdirAll(dir, 0755)
		f, err := os.Create(fullPath)
		if err != nil {
			log.Fatal(err)
		}
		_, err = io.Copy(f, tr)
		if err != nil {
			f.Close()
			return err
		}
		f.Chmod(info.Mode())
		defer f.Close()
	}
}

func main() {
	// first param is the dest path
	unTar(".", os.Stdin)
}
