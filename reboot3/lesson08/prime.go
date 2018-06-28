package main

import (
	"fmt"
)

func main() {
next:
	for out := 2; out < 10; out++ {
		fmt.Printf("\n-- out:out=%d\n", out)
		// < out，不需要和自己进行比较。
		for in := 2; in < out; in++ {
			fmt.Printf(" - in :out=%d,in=%d, result=%d\n", out, in, out%in)
			if out%in == 0 {
				continue next // contine goto
			}
		}
		fmt.Printf("go  %d\n\n", out)
	}
	fmt.Println("finished")
}
