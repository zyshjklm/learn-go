// for page 97 
package main 

import (
	"fmt"
)

// global map
var m = make(map[string]int)

// list to string
func toStr(list []string) string {
	return fmt.Sprintf("%q", list)
}

func Add(list []string) {
	m[toStr(list)]++
}

func Count(list []string) int {
	return m[toStr(list)]
}

var months[13]string = [...]string{
	1: "January",
	2: "February",
	3: "March",
	4: "April",
	5: "May",
	6: "June",
	7: "July",
	8: "August",
	9: "September",
	10: "October",
	11: "November",
	12: "December",
}

func main() {
	fmt.Println("before any:\n", m)

	Q2 := months[4:7]
	Add(Q2)
	fmt.Println("after add Q2:\n", m)
	fmt.Printf("count: %d\n", Count(Q2))

	summer := months[6:9]
	Add(summer)
	fmt.Println("after add summer:\n", m)
	fmt.Printf("count: %d\n", Count(summer))
}

