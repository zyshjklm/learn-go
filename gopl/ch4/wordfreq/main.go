// exercise 4.9
package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	counts := make(map[string]int)	
	
	input := bufio.NewScanner(os.Stdin)
	// break the input into words instead of lines.
	input.Split(bufio.ScanWords)

	for input.Scan() {
		word := input.Text()
		fmt.Println("\tloops: ", word)
		counts[word]++
	}

	// print after ctrl-D
	fmt.Println(counts)
}