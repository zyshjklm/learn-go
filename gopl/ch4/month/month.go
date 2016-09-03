// for 4.2 slices
package main

import (
    "fmt"
)

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
	Q2 := months[4:7]
	summer := months[6:9]

	fmt.Println(Q2)
	fmt.Println(summer)
	// [April May June]
	// [June July August]

	// find common element
	for _, j := range Q2 {
		for _, k := range summer {
			if j == k {
				fmt.Printf("%s appears in both\n", j)
			}
		}
	}

	// len, cap
	fmt.Printf("len of summer: %d\n", len(summer))
	fmt.Printf("cap of summer: %d\n", cap(summer))

	// summer[:20] will panic: out of range
	endlessSummer := summer[:5]		// within cap
	fmt.Println(endlessSummer)
}
