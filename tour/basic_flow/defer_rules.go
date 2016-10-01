package main 

import (
	"fmt"
)

// rules 1:
// deferred function's argument are evaluated when
// the defer statement is evalutead
func defer_rule1() int {
	i := 0
	// i in defer args is 0
	defer fmt.Println("\tin func1:", i)
	i++ 
	return i
}

// rule 2:
// deferred function calls are executed in Last-in-First-out
// order after the surrounding function returns
func defer_rule2() {
	for i := 0; i < 4; i++ {
		defer fmt.Print(i)
	}
}

// rule 3:
// deferred functions may read and assign to the returning
// function's named return values.
func defer_rule3(i int) int {
	defer func() { 
		i++; fmt.Println("\tin func3:", i)
	}()
	// i plus in defer after surrounding func return 1.
	return i
}

func main() {
	fmt.Printf("in main() call rule1: %d\n",  defer_rule1())
	
	fmt.Printf("in main() call rule2: ")
	defer_rule2()
	fmt.Println()

	fmt.Printf("in main() call rule3: %d\n",  defer_rule3(1))

}
