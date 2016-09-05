package main

import (
	"fmt"
)

type tree struct {
	value int
	left, right *tree
}

// add a value to a tree
func add(t *tree, value int) *tree {
	// fmt.Printf("add: start : %v\n", t)

	if t == nil {
		t = new(tree)
		t.value = value
		// fmt.Println("add: after create: ", t.value)
		return t
	}

	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}

	return t
}

// append the elements of t to values in order
// add returns the resulting slice.
func appendValue(value []int, t *tree) []int {
	if t != nil {
		value = appendValue(value, t.left)
		value = append(value, t.value)
		value = appendValue(value, t.right)
	}	

	return value
}

func Sort(value []int) {
	var root *tree

	for _, v := range value {
		root = add(root, v)
		fmt.Printf("Sort loops: %d\t%v\n", v, root)
	}	
	// write back to value slice
	appendValue(value[:0], root)
}

func main() {
	var base []int = []int{5, 8 ,7, 1, 6, 4}
	fmt.Println(base)

	Sort(base)
	fmt.Println(base)
}

