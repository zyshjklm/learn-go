package main

import (
	"bufio"
	"fmt"
	"os"
)


type DupLines struct {
	LineCount int
	FilesName []string
}

// Dup2_mod for exercise 1.4
// prints the names of all files in which each dumplicated lines occurs

func dup2_mod() {
	counts := make(map[string]*DupLines)
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

	for line, dupLine := range counts {
		// fmt.Println(line, dupLine)
		if dupLine.LineCount > 1 {
			fmt.Printf("%d\t%s\t%v\n", dupLine.LineCount, line, dupLine.FilesName)
		}
	}
}

// public func for dup2
func countLines(f *os.File, counts map[string]*DupLines) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		key := input.Text()

		_, ok := counts[key]
		if ok {
			counts[key].LineCount++
			// counts[key].FilesName = append(counts[key].FilesName, f.Name())
			counts[key].FilesName = updateArray(counts[key].FilesName, f.Name())
		} else {
			counts[key] = new(DupLines)
			counts[key].LineCount = 1

			// counts[key].FilesName = append(counts[key].FilesName, f.Name())
			counts[key].FilesName = updateArray(counts[key].FilesName, f.Name())
			
		}
	}
}


// prevent the given array from duplicating the element of file names.
// if f in array, then return directly.

func updateArray(filesName []string, f string) []string {
	for _, fn := range filesName {
		if fn == f {
			return filesName
		} 
	}
	filesName = append(filesName, f)
	return filesName

}

func main() {
	dup2_mod()
}


