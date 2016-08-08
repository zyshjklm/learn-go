// basename removes directory components and a .suffix.
package main 

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)



func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		fmt.Println(basename2(input.Text()))
	}	
	// NOTE: ignoring potential errors from input.Err()
}

// basename removes directory components and a .suffix.
// e.g.
// a => a; a.go => a; a/b/c.go => c; a/b.c.go => b.c
func basename1(s string) string {
	// Discard last '/' and everything before.
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}

	// preserve everthing before last '.'
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	//fmt.Println("before return:", s)
	return s
}

func basename2(s string) string {
	slash := strings.LastIndex(s, "/")	// -1 if not found
	s = s[slash+1:]

	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}
