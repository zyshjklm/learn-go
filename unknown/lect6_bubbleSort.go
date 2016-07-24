package main 

import "fmt"

func BubbleSort() {
	a := [...]int {5, 2, 6, 3, 9, 7, 1}
	num := len(a)

	fmt.Println("i j", a, "\n")

	for i := 0; i < num; i++ {
		for j := i+1; j < num-1; j++ {
			if a[i] < a[j] {
				a[i], a[j] = a[j], a[i]
			}
			fmt.Println(i, j, a)
		}
	}
	fmt.Println(a)
}

func main() {
	BubbleSort()
}



