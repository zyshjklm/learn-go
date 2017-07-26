package main

import (
	"fmt"
	"path"
	"path/filepath"
)

func main() {
	s := "/51reboot/golang/blob/master/lesson7/check.md"
	dir := path.Dir(s)
	name := path.Base(s)
	fmt.Println(dir, name)
	// /51reboot/golang/blob/master/lesson7 check.md

	// for win
	s1 := "\\51reboot\\golang\\blob\\master\\lesson7\\check.md"
	dir = path.Dir(s1)
	name = path.Base(s1)
	fullName := filepath.Join(dir, name)
	fmt.Println(dir, name)
	// . \51reboot\golang\blob\master\lesson7\check.md

	fmt.Println(fullName)
	// \51reboot\golang\blob\master\lesson7\check.md
}
