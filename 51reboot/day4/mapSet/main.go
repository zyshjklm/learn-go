package main

import "fmt"

func main() {
	sets := make(map[string]bool)
	sets["a"] = true
	dic := [2]string{"a", "b"}
	for idx := 0; idx < len(dic); idx = idx + 1 {
		// fmt.Printf("val:%v\n", dic[idx])
		if sets[dic[idx]] {
			fmt.Println("have it :", dic[idx])
		} else {
			fmt.Println("not have:", dic[idx])
		}
	}
}
