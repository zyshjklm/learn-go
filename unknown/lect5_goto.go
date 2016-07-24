package main 

import (
	"fmt"
)


func main() {

LABEL1:
	for {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
			if i > 3 {		// 0, 1, 2, 3. 4 to break
				break LABEL1
			}
		}
	}
	// break to here
	fmt.Println("OK LABEL1\n")


LABEL2:
	// continue to here
	for i := 0; i < 10; i++ {
		for {
			fmt.Println(i)
			continue LABEL2		// run 0-9 times
		}
	}
	fmt.Println("OK LABEL2\n")


	for {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
			if i > 3 {
				goto LABEL_GOTO
			}
		}
	}

LABEL_GOTO:
	fmt.Println("OK LABEL_GOTO\n")

}

