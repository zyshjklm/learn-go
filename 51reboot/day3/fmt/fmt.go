package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Create("fmt.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprint(f, "hello")
	fmt.Fprintln(f, "helloln")

	s := "hello"
	n := 4
	fmt.Fprintf(f, "my string is : %s n=%d\n", s, n)

	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			if j != 1 {
				fmt.Fprintf(f, " %d*%d=%d", i, j, i*j)
			} else {
				fmt.Fprintf(f, "%d*%d=%d", i, j, i*j)
			}

		}
		fmt.Fprintf(f, "\n")
	}
	f.Close()
}
