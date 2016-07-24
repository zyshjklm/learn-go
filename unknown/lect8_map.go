package main 

import "fmt"
import "sort"

// map
/*
	速度：
	索引 > map > 线性搜索

*/

func main() {
	// map[keyType]valuType
	// style 1
	var m0 map[int]string
	m0 = map[int]string{}
	fmt.Println(m0)
	
	// style 2
	var m1 map[int]string
	m1 = make(map[int]string)
	fmt.Println(m1)

	// style 3
	m2 := make(map[int]string)
	fmt.Println(m2)

	// use map by key:v
	m2[1] = "ok"
	fmt.Println(m2)
	a := m2[1]
	fmt.Println(a)

	delete(m2, 1)	// map, key
	fmt.Println(m2)

	// map nested
	var m3 map[int]map[int]string
	m3 = make(map[int]map[int]string)

	m3[1] = make(map[int]string)
	m3[1][1] = "OK"
	a2 := m3[1][1]
	fmt.Println(a2)

	a3, flag := m3[2][1]
	if !flag {
		m3[2] = make(map[int]string)
	}
	m3[2][1] = "GOOD"
	a3, flag = m3[2][1]
	fmt.Println(a3, flag)

	fmt.Println("\n--- for range ---\n")
	sm := make([]map[int]string, 5)	// length 5

	// for index, value := range slice
	// for key, value := range map
	
	for _, v := range sm {
		v = make(map[int]string, 1)
		v[1] = "OK"
		fmt.Println(v)
	}
	// v just a copy, not effect sm
	fmt.Println(sm, "\n")

	for i := range sm {
		sm[i] = make(map[int]string, 1)
		sm[i][1] = "OK"
		fmt.Println(i, sm[i])
	}
	// i effect sm
	fmt.Println(sm)

	map1 := map[int]string{1:"a", 2: "b", 3:"c", 4:"d"}
	sls1 := make([]int, len(map1))
	i := 0
	for k, _ := range map1 {
		sls1[i] = k
		i++
	}
	fmt.Println(sls1)
	sort.Ints(sls1)
	fmt.Println(sls1)
}


