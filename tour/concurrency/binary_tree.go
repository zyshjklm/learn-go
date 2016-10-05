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

func main() {
	Walk_test()
}