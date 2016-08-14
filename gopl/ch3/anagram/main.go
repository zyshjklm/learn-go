// exercise 3.11. anagram

package main 

import (
	"fmt"

)

func anagram(s1, s2 string) bool {
	leterCount1 := map[string]int{}
	leterCount2 := map[string]int{}
	// fmt.Println(leterCount1, leterCount2)

	for _, v := range s1 {
		leterCount1[string(v)] += 1
	}
	for _, v := range s2 {
		leterCount2[string(v)] += 1
	}

	fmt.Println(leterCount1, leterCount2)

	for k, v := range leterCount1 {
		if v != leterCount2[k] {
			return false
		}
	}
	return true
}

func main() {
	// test 
	fmt.Println(anagram("tops", "post"))
	fmt.Println(anagram("tops", "poss"))	
}