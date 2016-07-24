package main 

import "fmt"

func main() {
	map1 := map[int]string{4:"d",1:"a", 2: "b", 3:"c", }
	map2 := make(map[string]int)

	// reverse map
	for k, v := range map1 {
		fmt.Println(k, v)
		map2[v] = k
	}
	fmt.Println(map1)
	fmt.Println(map2)

	map3 := map[string]string{"11":"aa", "22": "bb", "33":"cc", "44":"dd"}
	map4 := make(map[string]string)

	for k, v := range map3 {
		fmt.Println(k, v)
		map4[v] = k
	}
	fmt.Println(map3)
	fmt.Println(map4)	

	i := 0
	sls0 := make([]int, len(map1))
	for k, v := range map1 {
		fmt.Println(k, v)
		sls0[i] = k
		i++
	}
	fmt.Println(map1)
	fmt.Println(sls0)
}


