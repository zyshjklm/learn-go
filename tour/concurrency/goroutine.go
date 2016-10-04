package main 

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100*time.Millisecond)
		fmt.Println(s)
	}
}

func no_rountine() {
	say("world")
	say("Hello")
}

func has_rountine() {
	go say("world")
	say("Hello")
}

func main() {
	has_rountine()
	// pls diff with no_rountine()
}