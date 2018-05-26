package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []string{"Go", "Bravo", "Gohper", "Alpha", "Grin", "Delta"}
	fmt.Println(s)
	sort.Strings(s)
	fmt.Println(s)
}
