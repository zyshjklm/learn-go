package main

import (
	"fmt"

	. "github.com/jkak/learn-go/reboot3/lesson09/pkgbase/a"
	. "github.com/jkak/learn-go/reboot3/lesson09/pkgbase/b"
)

func init() {
	fmt.Println("I'm main()")
}

func main() {
	A()
	B()
}
