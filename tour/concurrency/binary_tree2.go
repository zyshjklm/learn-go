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



func main() {
	Walk_test()
	
}
