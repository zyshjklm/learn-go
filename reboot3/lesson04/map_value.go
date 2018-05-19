package main

import "fmt"

func main() {
	ages := make(map[string]int)
	ages["tome"] = 23
	ages["jack"] = 24

	if v, ok := ages["jack"]; ok {
		fmt.Println(v)
	}

	if v, ok := ages["jack2"]; ok {
		fmt.Println(v)
	}
	for k, v := range ages {
		fmt.Printf("%s -> %d\n", k, v)
	}
}
