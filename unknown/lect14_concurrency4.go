package main 

import (
	"fmt"
)

func pingPong(ch chan string) {
	var i int64 = 10000

	for {
		// read from channel	
		fmt.Println("routine read ch:", <- ch)

		// write
		ch <- fmt.Sprintf("Hi, #%d", i)
		i++
	}
}

func main() {
	chStr := make(chan string)

	go pingPong(chStr)

	for i := 0; i < 10; i++ {
		// write
		chStr <- fmt.Sprintf("Hello, #%d", i)
		
		// read from channel	
		fmt.Println(" - main read ch:", <- chStr)	
	}

}


