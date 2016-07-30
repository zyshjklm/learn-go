package main 

import (
	"fmt"
	"time"
)


func main() {

	s1 := []string{"a", "b", "c"}

	// pits of closure
	for _, v := range s1 {
		go func() {
			fmt.Println(v)	// c for each routine
		}()
	}

	time.Sleep(time.Second)
	fmt.Println()

	// fix it
	for _, v := range s1 {
		go func(v string) {
			fmt.Println(v)	
		}(v)
	}
	time.Sleep(time.Second)
	// select{}

}