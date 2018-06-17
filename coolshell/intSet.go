package main

import (
	"fmt"
	"sort"
	"strings"
)

// IntSet for int set
type IntSet struct {
	data map[int]bool
}

// NewIntSet to new an IntSet
func NewIntSet() IntSet {
	return IntSet{make(map[int]bool)}
}

// Add to add an int
func (set *IntSet) Add(x int) {
	set.data[x] = true
}

// Delete to delete an int
func (set *IntSet) Delete(x int) {
	delete(set.data, x)
}

// Has to judge x in a Set
func (set *IntSet) Has(x int) bool {
	return set.data[x]
}

// String to satisfies fmt.Stringer interface
func (set *IntSet) String() string {
	if len(set.data) == 0 {
		return "{}"
	}
	ints := make([]int, 0, len(set.data))
	for i := range set.data {
		ints = append(ints, i)
	}
	sort.Ints(ints)

	parts := make([]string, 0, len(ints))
	for _, i := range ints {
		parts = append(parts, fmt.Sprint(i))
	}
	return "{" + strings.Join(parts, ",") + "}"
}

func main() {
	ints := NewIntSet()
	for _, i := range []int{1, 3, 5, 7} {
		ints.Add(i)
		fmt.Println(ints)
	}

	for _, i := range []int{1, 2, 3, 4, 5, 6, 7} {
		fmt.Print(i, ints.Has(i), "\t")
		ints.Delete(i)
		fmt.Println(ints)
	}
	// {map[1:true]}
	// {map[1:true 3:true]}
	// {map[5:true 1:true 3:true]}
	// {map[1:true 3:true 5:true 7:true]}
	// 1 true	{map[3:true 5:true 7:true]}
	// 2 false	{map[3:true 5:true 7:true]}
	// 3 true	{map[5:true 7:true]}
	// 4 false	{map[5:true 7:true]}
	// 5 true	{map[7:true]}
	// 6 false	{map[7:true]}
	// 7 true	{map[]}
}
