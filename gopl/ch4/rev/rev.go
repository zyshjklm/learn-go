package main

import "fmt"

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5}
	fmt.Println(a)

	reverse(a[:])
	fmt.Println(a)
	reverse(a[:])
	fmt.Println("# before rotate left:\n", a)

	// rotate left. reverse all finally.
	fmt.Println("\n--- start to rotate s left by 2 ---\n")
	reverse(a[:2])
	fmt.Println(a)	
	reverse(a[2:])
	fmt.Println(a)
	
	reverse(a[:])
	fmt.Println(a)

	// rotate right. reverse all first.
	a = [...]int{0, 1, 2, 3, 4, 5}
	fmt.Println("\n# before rotate right:\n", a)

	fmt.Println("\n--- start to rotate s right by 2 ---\n")
	reverse(a[:])
	fmt.Println(a)

	reverse(a[:2])
	fmt.Println(a)	
	reverse(a[2:])
	fmt.Println(a)
	
}