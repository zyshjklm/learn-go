package main

import (
	"fmt"
	"os"
	"strconv"
)

func add(m, n int) int {
	return m + n
}

func sub(m, n int) int {
	return m - n
}

type counter func(int, int) int

func main() {
	funcmap := map[string]counter{
		"+": add,
		"-": sub,
	}
	funcmap["+"](13, 5)
	funcmap["-"](13, 5)

	m, _ := strconv.Atoi(os.Args[1])
	n, _ := strconv.Atoi(os.Args[3])
	f := funcmap[os.Args[2]]
	if f != nil {
		fmt.Println(f(m, n))
	}
}
