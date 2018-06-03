package main

import "fmt"

func main() {
	s := "hello golang"
	arr := []byte(s)
	arr[2] = 'w'
	fmt.Println(string(arr))
}
