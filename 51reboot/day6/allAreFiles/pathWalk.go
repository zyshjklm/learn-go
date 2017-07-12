package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// filepath.Walk() and it WalkFunc, refer:
// http://localhost:6060/pkg/path/filepath/#Walk

// type WalkFunc func(path string, info os.FileInfo, err error) error
func walkFn(path string, info os.FileInfo, err error) error {
	fmt.Println("path name:", path)
	if info.IsDir() {
		fmt.Println("-- dir:", path)
	}
	return nil
}

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("usage: %s dir/path\n", os.Args[0])
		os.Exit(1)
	}

	dir := os.Args[1]
	filepath.Walk(dir, walkFn)
}
