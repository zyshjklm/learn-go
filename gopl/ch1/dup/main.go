package main

import (
	"bufio"
	"fmt"
	"os"

	"io/ioutil"
	"strings"
)



// Dup1 prints the text of each line that appears more than
// once in the standard input, preceded by its count.

func dup1() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	// end by CTRL+D 

	for input.Scan() {
		counts[input.Text()]++
	}

	// use Println() or Printf("\n") to output ^D before result. 
	fmt.Println()	

	// NOTE: ignoring potential errors from input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}


// Dup2 prints the count and text of lines that appears more than
// once in the input. It reads from stdin or from a list of named files.

func dup2() {
	counts := make(map[string]int)
	files := os.Args[1:]

	// read from stdio or files
	if len(files) == 0 {
		countLines(os.Stdin, counts)
		// use Println() or Printf("\n") to output ^D before result. 
		fmt.Println()
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

// public func for dup2
func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
}


// dup3 prints the count and text of ines that
// appear more than once ni the named input files.

func dup3() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}

		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func main() {
	// dup1()
	dup2()
	// dup3()
}


