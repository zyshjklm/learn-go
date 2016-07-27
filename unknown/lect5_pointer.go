package main 

import (
	"fmt"
)


func main() {
	var x int = 65
	var p *int = &x		// &x 取地址

	fmt.Println(p)
	fmt.Println(*p)		// *p 取值
	
}

