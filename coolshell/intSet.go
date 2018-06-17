package main

import (
	"errors"
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

// --- for Undo ---

// UndoableIntSet as an undoable IntSet
type UndoableIntSet struct {
	IntSet // embedding(delegation)
	Funcs  []func()
}

// NewUndoableIntSet to new an NewUndoableIntSet
func NewUndoableIntSet() UndoableIntSet {
	return UndoableIntSet{NewIntSet(), nil}
}

// Add to add an int to UndoableIntSet
func (set *UndoableIntSet) Add(x int) {
	// override Add() of IntSet
	if !set.Has(x) {
		set.data[x] = true
		set.Funcs = append(set.Funcs, func() { set.Delete(x) })
	} else {
		set.Funcs = append(set.Funcs, nil)
	}
}

// Delete to delete an int from UndoableIntSet
func (set *UndoableIntSet) Delete(x int) {
	// override Delete() of IntSet
	if set.Has(x) {
		delete(set.data, x)
		set.Funcs = append(set.Funcs, func() { set.Add(x) })
	} else {
		set.Funcs = append(set.Funcs, nil)
	}
}

// Undo to undo an Add or Delete ops
func (set *UndoableIntSet) Undo() error {
	if len(set.Funcs) == 0 {
		return errors.New("no funcs to undo")
	}
	index := len(set.Funcs) - 1

	if fName := set.Funcs[index]; fName != nil {
		fName()
		set.Funcs[index] = nil // free closure for barbage collection
	}
	fmt.Printf("- undo index %3d: %+v\n", index, set.IntSet)
	set.Funcs = set.Funcs[:index]
	return nil
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

	fmt.Println("\n---- Undo test ----")
	undoInts := NewUndoableIntSet()
	for _, i := range []int{1, 3, 5, 7} {
		undoInts.Add(i)
		fmt.Printf("%v\n\tundo funcs len:%d\n", undoInts.IntSet, len(undoInts.Funcs))
	}

	fmt.Println("\n-- Delete() --")
	for _, i := range []int{1, 2, 3, 4, 5, 6, 7} {
		fmt.Print(i, undoInts.Has(i), "\t")
		undoInts.Delete(i)
		fmt.Printf("%v\n\tfuncs len:%3d\n",
			undoInts.IntSet, len(undoInts.Funcs))
	}

	fmt.Println("\n-- Undo() --")
	for {
		if err := undoInts.Undo(); err != nil {
			break
		}
	}
	fmt.Println(undoInts)
}
