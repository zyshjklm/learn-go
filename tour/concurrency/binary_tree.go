package main 

import (
	"fmt"
	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

func Walk_test() {
	var t *tree.Tree = tree.New(1)
	fmt.Println(t)
	
	// tree.New() generate 10 Values.
	var ch = make(chan int, 10)
	Walk(t, ch)
	close(ch)

	for i := range ch {
		fmt.Printf("%d ", i)
	}
}

// determines whether the trees 
// t1 an t2 contains the same values.
func Same(t1, t2 *tree.Tree) bool {
	var ch1 = make(chan int, 10)
	var ch2 = make(chan int, 10)
	Walk(t1, ch1)
	Walk(t2, ch2)
	close(ch1)
	close(ch2)

	var c1, c2 int
	for c1 = range ch1 {
		c2 = <- ch2
		fmt.Printf("ch1: %d, ch2: %d\n", c1, c2)
		if c1 != c2 { return false }
	}
	return true
}


func main() {
	// Walk_test()

	var t1 *tree.Tree = tree.New(1)
	fmt.Println(t1)
	var t2 *tree.Tree = tree.New(1)
	fmt.Println(t2)
	
	var t3 *tree.Tree = tree.New(3)
	fmt.Println(t3)

	fmt.Printf("\nis t1 same to t2? %v\n", Same(t1, t2))	
	fmt.Printf("\nis t1 same to t3? %v\n", Same(t1, t3))	
}
