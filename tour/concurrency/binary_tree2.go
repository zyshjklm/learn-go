package main 

import (
	"fmt"
	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return 
	}
	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
}

func Walk_test() {
	var t *tree.Tree = tree.New(1)
	fmt.Println(t)
	
	var ch = make(chan int)

	go func() {
		Walk(t, ch) 
		close(ch)
	}()

	for i := range ch {
		fmt.Printf("%d ", i)
	}	
	fmt.Println("\n")	
}

// determines whether the trees 
// t1 an t2 contains the same values.
func Same(t1, t2 *tree.Tree) bool {
	var ch1, ch2 = make(chan int), make(chan int)

	go func() {
		Walk(t1, ch1)
		close(ch1)
	}()

	go func() {
		Walk(t2, ch2)
		close(ch2)
	}()

	for {
		c1, ok1 := <- ch1
		c2, ok2 := <- ch2
		if ok1 != true || ok2 != true {
			break	// channel closed
		}
		fmt.Printf("  ch1: %d, ch2: %d\n", c1, c2)
		if c1 != c2 { 
			return false 
		}
	}
	return true
}


func main() {
	Walk_test()
	
	var t1, t2, t3 *tree.Tree
	
	t1 = tree.New(1)
	t2 = tree.New(1)
	t3 = tree.New(2)
	fmt.Println("t1:", t1)
	fmt.Println("t2:", t2)
	fmt.Println("t3:", t3)

	fmt.Printf("is t1 same to t2? %v\n", Same(t1, t2))	
	fmt.Printf("is t1 same to t3? %v\n", Same(t1, t3))

}
