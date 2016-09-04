// exercise 4-5 : eliminate adjacent duplicates in a []string slice

package main 

import (
	"fmt"
)

func dupAdjacent(str []string) []string {
	var result []string
	result = str[0:1]
	i := 0

	for _, s := range str {
		if result[i] != s {
			result = append(result, s)
			i++
		}
	}
	return result
}

func dupAdjacent2(str []string) []string {
	result := make([]string, len(str), len(str))
	result[0] = str[0]
	i := 0

	for _, s := range str {
		if result[i] != s {
			i++
			result[i] = s
		}
	}
	return result[:i+1]
}

func main() {
	testStr := [...]string{"1", "1", "1", "2", "2", "3"}
	fmt.Println(testStr)
	newStr := dupAdjacent(testStr[:])
	fmt.Println(newStr)

	testStr2 := [...]string{"1", "1", "1", "2", "2", "3"}
	newStr2 := dupAdjacent2(testStr2[:])
	fmt.Println(newStr2)
}