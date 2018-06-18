package main

import "fmt"

func main() {
	// basic type
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}

	// classic initial/condition/after loop
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// without a condition
	for {
		fmt.Println("loop")
		break
	}

	// continue
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
