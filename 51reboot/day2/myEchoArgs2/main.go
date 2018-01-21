package main

import (
	"flag"
	"fmt"
	"strings"
)

var newline = flag.Bool("n", false, "end with \n\t for string")
var sep = flag.String("s", " ", "separator")
var myType string

// sep 的类型是推导出来的，是字符串指针
// -- help的结果：
// Usage of ./myecho1:
//  -s string
//        separator (default " ")

func main() {
	flag.StringVar(&myType, "t", "test", "test, prod, preview")
	flag.Parse()
	//fmt.Print(strings.Join(flag.Args(), *sep))

	str := strings.Join(flag.Args(), *sep)
	if *newline {
		fmt.Println(str)
	} else {
		fmt.Print(str)
	}
}
