package main 

import (
	"fmt"
	"time"
)


func main() {
	// 每隔多少秒
	tick := time.Tick(  400 * time.Millisecond)
	// sleep但不阻塞
	boom := time.After(1200 * time.Millisecond)

	for {
		select {
		case <- tick:
			fmt.Println("tick.")
		case <- boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(200 * time.Millisecond)
		}
	}
}
