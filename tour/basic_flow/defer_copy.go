package main 

import (
	"fmt"
	"os"
	"io"
)

func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()

	dst, err := os.Create(dstName)
	if err != nil {
		return
	}
	defer dst.Close()

	return io.Copy(dst, src)
}

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("usage: %s dstFile srcFile\n", os.Args[0])
		os.Exit(1)
	}
	CopyFile(os.Args[1], os.Args[2])
}
