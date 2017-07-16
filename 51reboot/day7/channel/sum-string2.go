package main

import "fmt"

func sum(s []string, c chan string) {
	sum := ""
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	s := []string{"hello", "world", "golang", "c++", "python"}
	c1 := make(chan string)
	c2 := make(chan string)

	half := len(s) / 2
	fmt.Println("\nhalf:", half)
	go sum(s[:half], c1)
	go sum(s[half:], c2)
	// 并行执行，但有确定的槽位，通过不能的通道来区分
	x, y := <-c1, <-c2 // receive from c

	fmt.Printf("%s\n%s\n%s\n", x, y, x+y)
}
