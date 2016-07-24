package main 

import "fmt"

func main() {

	// append to slice
	sl1 := make([]int, 3, 6)
	fmt.Printf("%p : %v\n", sl1, sl1)
	sl1 = append(sl1, 1, 2, 3)
	fmt.Printf("%p : %v\n", sl1, sl1)
	sl1 = append(sl1, 4, 5, 6)
	fmt.Printf("%p : %v\n", sl1, sl1)

	sl3 := make([]int, 3, 6)
	fmt.Println("\n", sl3, "\n")

	for i := 3; i < 15; i++ {
		if len(sl3) <= cap(sl3) {
			sl3 = append(sl3, i)
		}
		fmt.Println("i:", i, "len:",len(sl3), 
			"cap:",cap(sl3), sl3)
		fmt.Printf("  Addr: %p\n", sl3)
	}

	sl4 := [5]int{1,2,3,4,5}

	sl5 := sl4[2:5]
	sl6 := sl4[1:3]

	fmt.Println("\nsl4:", sl4, "\nsl5:", sl5, "\nsl6:", sl6)
	sl6[1] = 9
	fmt.Println("\nsl4:", sl4, "\nsl5:", sl5, "\nsl6:", sl6)
	// slice affect array

	fmt.Printf("addr of sl6: %p\n", sl6)
	sl6 = append(sl6, 6,7,8,9,0)
	fmt.Printf("addr of sl6: %p\n", sl6)	// new addr

	sl6[1] = 8
	fmt.Println("\nsl4:", sl4, "\nsl5:", sl5, "\nsl6:", sl6)
	
	// copy
	sl8 := []int{1,2,3,4,5,6}
	sl9 := []int{7,8,9}
	fmt.Println("\ncopy:\n")
	fmt.Println(sl8, sl9)
	copy(sl8, sl9)		// sl8[0:4] equal sl9
	fmt.Println(sl8, sl9)	

	sl10 := []int{11,22,33}
	fmt.Println(sl10, sl8)
	copy(sl10, sl8)		// sl10[0:4] equal sl8
	fmt.Println(sl10, sl8)

	copy(sl10, sl8[3:])
	fmt.Println(sl10, sl8)

	// copy use slice
	// [:], [:len(slice)]
	sl11 := sl8[:]
	fmt.Println(sl11, sl8)
}
