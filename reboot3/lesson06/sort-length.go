package main

import (
	"fmt"
	"sort"
)

type StringSlice []string

func (p StringSlice) Len() int           { return len(p) }
func (p StringSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p StringSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

type StringSlice2 []string

func (p StringSlice2) Len() int           { return len(p) }
func (p StringSlice2) Less(i, j int) bool { return len(p[i]) < len(p[j]) }
func (p StringSlice2) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func main() {
	s := []string{"ab", "B", "CCC"}

	sort.Strings(s)
	fmt.Println(s)

	sort.Sort(StringSlice(s))
	fmt.Println(s)

	sort.Sort(StringSlice2(s))
	fmt.Println(s)
}
