package main

import "fmt"

func main() {
	fmt.Println("start main()")
	panic("trigger panic in main.")
	fmt.Println("over!")
	// outout:
	// start main()
	// panic: trigger panic in main.
}
