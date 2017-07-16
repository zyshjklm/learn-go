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

	c := make(chan string)
	half := len(s) / 2
	fmt.Println("half:", half)
	go sum(s[:half], c)
	go sum(s[half:], c)
	x, y := <-c, <-c // receive from c
	// 共用一个通道

	fmt.Printf("%s\n%s\n%s\n", x, y, x+y)
}
