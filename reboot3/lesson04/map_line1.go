package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var seen map[string]bool
	seen = map[string]bool{}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("> ")
		scanner.Scan()
		line := scanner.Text()
		if !seen[line] {
			fmt.Println(line)
			seen[line] = true
		}
	}
}
