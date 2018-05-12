package main

import "fmt"

func main() {
	a := [3]string{"one", "two", "three"}
	// reverse
	result := make([]string, 0, 3)
	for i := len(a) - 1; i >= 0; i-- {
		result = append(result, a[i])
	}
	fmt.Println(result)
}
