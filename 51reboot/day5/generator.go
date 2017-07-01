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
}
