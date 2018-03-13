package main

import (
	"archive/tar"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

func maketar(dir string, w io.Writer) error {
	baseDir := filepath.Base(dir)
	tr := tar.NewWriter(w)
	defer tr.Close()

	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		header, err := tar.FileInfoHeader(info, "")
		if err != nil {
			return err
		}
		// 注意：上面是通过info区创建一个tar的header。
		// 但该函数返回的header中的名字已经只是base name了
		fmt.Printf("\npath=%s, header.name=%s, info.name=%s\n", path, header.Name, info.Name())
		// header.Name = path
		// 计算当前循环的路径path相对于基础路径dir的相对路径。
		relPath, _ := filepath.Rel(dir, path)
		fmt.Println("relative path:", relPath)
		header.Name = filepath.Join(baseDir, relPath)
		tr.WriteHeader(header)

		f, err := os.Open(path)
		if err != nil {
			return err
		}
		defer f.Close()
		io.Copy(tr, f)
		return nil
	})
	return nil
}

func main() {
	dstName := "img.tar"
	fd, err := os.Create(dstName)
	if err != nil {
		log.Fatal(err)
	}
	maketar("./", fd)
}
