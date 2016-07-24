package main 

import "fmt"

func funcA(a int, b string) {
	fmt.Println(a, b)
}

func funcB(a, b, c int) int {
	fmt.Println(a, b, c)
	return a + b + c
}

// return 
func funcC() (a, b, c int) {
	a, b, c = 1, 2, 3
	return 
	// return a, b, c is recomment.
}

func funcD() (int, int, int) {
	a, b, c := 4, 5, 6 
	return a, b, c
}

// p 不定长参数，须放在参数列表的最后
func funcE(p ...int) {
	fmt.Println(p)
}

func funcF(s []int) {
	fmt.Println("-- modify slice...")
	s[0] = 1
	s[1] = 2
	s[2] = 3
	fmt.Println(s)
}

func funcG() {
	fmt.Println("in funcG...")
}

func main() {
	funcA(12, "34")
	fmt.Println(funcB(1, 2, 3))
	fmt.Println(funcC())

	fmt.Println(funcD())

	funcE(1,2,3,4)
	funcE(6, 7, 8, 9, 0)

	fmt.Println("\n--- origin slice...")
	slice := []int{6, 7, 8, 9, 0}
	fmt.Println(slice)
	funcF(slice)
	fmt.Println(slice)

	newFunc := funcG
	newFunc()
}


