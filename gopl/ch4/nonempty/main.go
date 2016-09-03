package	main 

// nonempty returns a slice holding only the non-empty strings.
// the underlying array is modified during the call.

import (
	"fmt"
)

func nonempty(strs []string) []string {
	i := 0
	for _, s := range strs {
		if s != "" {
			strs[i] = s
			i++
		}
	}
	return strs[:i]
}

func main() {
	data := []string{"one", "", "three"}

	fmt.Printf("%q\n", nonempty(data))
	// ["one" "three"]
	fmt.Printf("%q\n", data)
	// ["one" "three" "three"]

	data1 := []string{"one", "", "", "three", "four"}

	data1 = nonempty(data1)
	fmt.Printf("%q\n", data1)
	// ["one" "three", "four"]
}
