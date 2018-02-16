package main

import (
	"errors"
	"fmt"
)

func iter(s []int) func() (int, error) {
	var i int
	return func() (int, error) {
		if i >= len(s) {
			return 0, errors.New("end")
		}
		i++
		return s[i-1], nil
	}
}

func main() {
	s := []int{1, 2, 3, 4}
	f := iter(s)
	for {
		n, err := f()
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(n)
	}
	fmt.Print("\n--- style 2 ---\n")
	end := 10
	for i := 1; i < end; i++ {
		f := iter([]int{i})
		n, err := f()
		if err != nil {
			break
		}
		fmt.Println(n)
	}
}
