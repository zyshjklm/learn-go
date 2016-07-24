package main 

import (
	"fmt"
)


func main() {
	a := 10

	// if
	if a := 1; a > 0 {	// a is local var of if
		fmt.Println(a)	// 1
	}
	fmt.Println(a, "\n")	// 10

	// for 
	a = 0	// assign value
	for {	// 无条件，死循环。需要内部判断时机做跳出
		a++
		if a > 3 {
			break
		}
		fmt.Println(a)	// 1, 2, 3
	}
	fmt.Println("break for over\n")

	// a is 4 now. a++: 5, 6, 7
	for i := 0; i < 3; i++ {
		a++
		fmt.Println(a)
	}
	fmt.Println("a++ over\n")

	i := 1
	switch i {
	case 0:
		fmt.Println("i=0")
	case 1:
		fmt.Println("i=1")
	default:
		fmt.Println("None")
	}

	fmt.Println("\n-- switch 2 --")
	fmt.Println("before switch, a =", a)
	switch {
	case a >= 0:
		fmt.Println("a>=0")
		fallthrough
	case a >= 1:
		fmt.Println("a=1")
	default:
		fmt.Println("None")
	}

	fmt.Println("\n-- switch 3 --")
	switch b := 2; {	// 赋值后要有;
	case b >= 0:
		fmt.Println("b>=0")
		fallthrough
	case b >= 1:
		fmt.Println("b>=1")
		fallthrough
	case b >= 2:
		fmt.Println("b>=2")
	default:
		fmt.Println("None")
	}
}

