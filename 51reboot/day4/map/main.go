package main

import "fmt"

func main() {
	ages1 := make(map[string]int)
	ages1["a"] = 1
	ages1["b"] = 2

	ages2 := map[string]int{
		"x": 11,
		"y": 12,
	}

	fmt.Println(ages1)
	fmt.Println(ages2)

	fmt.Println("-- get value from map, judge ok --")
	c, ok := ages1["c"]
	if !ok {
		fmt.Println("not found, and value=", c)
	} else {
		fmt.Println(c)
	}

	if c, ok := ages2["c"]; ok {
		fmt.Println(c)
	}
	ages := map[string]int{
		"joke":   3,
		"jungle": 5,
	}

	fmt.Println("-- range map --")
	for name, age := range ages {
		fmt.Println("name", name, "age", age)
	}
	// range key
	for name := range ages {
		fmt.Println(name)
	}
}
