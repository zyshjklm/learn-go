package main

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

// FuncInfo for func info
type FuncInfo struct {
	Func func()
	Name string
}

// NewFuncInfo to new FuncInfo
func NewFuncInfo(f func(), name string) FuncInfo {
	return FuncInfo{f, name}
}

// Undo as slice of FuncInfo
type Undo []FuncInfo

// Add for Undo
func (u *Undo) Add(f func(), name string) {
	fi := NewFuncInfo(f, name)
	*u = append(*u, fi)
}

// Undo for Undo type
func (u *Undo) Undo() error {
	bakUndo := *u
	if len(bakUndo) == 0 {
		return errors.New("no funcs to undo")
	}
	index := len(bakUndo) - 1
	fi := bakUndo[index]
	if fName := fi.Func; fName != nil {
		fName()
		bakUndo[index].Func = nil // free closure for barbage collection
	}
	fmt.Printf("- undo index %3d: %s\t", index, fi.Name)
	*u = bakUndo[:index]
	return nil
}

// IntSet for int set
type IntSet struct {
	data map[int]bool
	undo Undo
}

// NewIntSet to new an IntSet
func NewIntSet() IntSet {
	return IntSet{data: make(map[int]bool), undo: nil}
}

// Has to judge x in a Set
func (set *IntSet) Has(x int) bool {
	return set.data[x]
}

// Add to add an int
func (set *IntSet) Add(x int) {
	if !set.Has(x) {
		set.data[x] = true
		set.undo.Add(func() { set.Delete(x) }, "add")
	} else {
		set.undo.Add(nil, "nil")
	}
}

// Delete to delete an int
func (set *IntSet) Delete(x int) {
	if set.Has(x) {
		delete(set.data, x)
		set.undo.Add(func() { set.Add(x) }, "del")
	} else {
		set.undo.Add(nil, "nil")
	}
}

// Undo for set
func (set *IntSet) Undo() error {
	return set.undo.Undo()
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
	undoInts := NewIntSet()
	for _, i := range []int{1, 3, 5, 7} {
		undoInts.Add(i)
		undoLen := len(undoInts.undo)
		fmt.Printf("%v\n\tundo funcs len:%d : %+v\n",
			undoInts.data, undoLen, undoInts.undo[undoLen-1].Name)
	}

	fmt.Println("\n-- Delete() --")
	for _, i := range []int{1, 2, 3, 4, 5, 6, 7} {
		fmt.Print(i, undoInts.Has(i), "\t")
		undoInts.Delete(i)
		undoLen := len(undoInts.undo)
		fmt.Printf("%v\n\tfuncs len:%3d : %s\n",
			undoInts.data, undoLen, undoInts.undo[undoLen-1].Name)
	}

	fmt.Println("\n-- Undo() --")
	for {
		if err := undoInts.Undo(); err != nil {
			break
		}
		fmt.Println(undoInts.data)
	}
}
