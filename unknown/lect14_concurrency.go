package main 

import (
	"fmt"
//	"time"
)

func testGo() {
	fmt.Println("Go Go Go!!!")
}


func main() {
	// trick style: use time.Sleep
	go testGo()
	//time.Sleep(1 * time.Second)	

	// use channel
	cnl := make(chan bool)
	go func() {
		fmt.Println("Go Go Go by channel!!!")
		cnl <- true
	}()
	<-cnl

	
	// for 
	go func() {
		fmt.Println("Go Go Go by channel!!!")
		cnl <- true
		close(cnl)
	}()
	for v := range cnl {
		fmt.Println(v)
	}
}