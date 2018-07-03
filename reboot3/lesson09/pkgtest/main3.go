package main

import (
	"fmt"

	_ "github.com/jkak/learn-go/reboot3/lesson09/pkgbase/a"
	_ "github.com/jkak/learn-go/reboot3/lesson09/pkgbase/b"
)

func init() {
	fmt.Println("I'm init in main")
}

func main() {
	fmt.Println("in main()")
}
