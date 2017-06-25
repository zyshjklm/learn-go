package main

import (
	"fmt"
	"reflect"
)

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println("type of primes:", reflect.TypeOf(primes))

	var s []int = primes[1:4]
	fmt.Println("type of s:", reflect.TypeOf(s))

	fmt.Println(s)
	fmt.Println(&s[0])
	fmt.Println(&primes[1])

	var s1 []int
	s1 = s

	fmt.Println(&s1[0] == &s[0])

	names := [4]string{"john", "Paul", "George", "Ringo"}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)

	names = [4]string{"a", "b", "c", "d"}
	a = names[0:2]
	b = names[1:3]
	fmt.Println(a, b)
	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
	c := a[1:2]
	c[0] = "YYY"
	fmt.Println(a, b)
	fmt.Println(names)

	fmt.Println("---")
	s2 := []int{2, 3, 5, 7, 9, 11, 13}
	printSlice(s2)
	fmt.Println(&s2[0])

	s2 = s2[:0] // len 0,
	printSlice(s2)
	fmt.Println(&s2)

	s2 = s2[:4]
	printSlice(s2)
	fmt.Println(&s2[0])

	s2 = s2[2:]
	printSlice(s2)
	fmt.Println(&s2[0])

	// nil
	var s3 []int
	fmt.Println(s3, len(s3), cap(s3))
	if s3 == nil {
		fmt.Println("nil!")
	}
}

func printSlice(s []int) {
	fmt.Printf("\nlen=%d cap=%d %v\n", len(s), cap(s), s)
}
