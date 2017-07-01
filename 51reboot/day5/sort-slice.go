package main

import (
	"fmt"
	"sort"
)

type Student struct {
	Name string
	Id   int
}

func main() {
	s := []int{2, 3, 1, 5, 9, 7}

	fmt.Println(s)
	sort.Slice(s, func(i, j int) bool {
		return s[i] > s[j]
	})
	fmt.Println(s)

	ss := []Student{}
	ss = append(ss, Student{
		Name: "aa",
		Id:   2,
	})
	ss = append(ss, Student{
		Name: "bb",
		Id:   3,
	})
	ss = append(ss, Student{
		Name: "cc",
		Id:   1,
	})

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Name < ss[j].Name
	})
	fmt.Println(ss)
	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Id < ss[j].Id
	})
	fmt.Println(ss)
}
