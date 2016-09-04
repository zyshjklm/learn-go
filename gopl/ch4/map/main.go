// for 4.3 map
package main 

import (
	"fmt"
	"sort"
)

func main() {
	ages := make( map[string]int )
	ages["alice"] = 32
	ages["charlie"] = 30
	ages["bob"] = 25

	fmt.Println(ages["charlie"])
	ages["charlie"]++
	fmt.Println(ages["charlie"])

	// random order for name
	for name, age := range ages {
		fmt.Printf("%s\t%d\n", name, age)
	}

	// sort key
	fmt.Println("-- sort by key --")
	var names []string
	for name := range ages {
		names = append(names, name)
	}
	sort.Strings(names)

	for _, name := range names {
		age, ok := ages[name]
		fmt.Printf("ok: %t\t%d\n", ok, age)
		fmt.Printf("%s\t%d\n", name, age)
	}
}