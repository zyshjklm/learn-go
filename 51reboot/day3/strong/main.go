package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	var n int
	var f float32
	// f = n / 3
	// error| cannot use n / 3 (type int) as type float32 in assignment
	n = 10
	f = float32(n / 3)
	fmt.Println(f, n)  // 3, 10
	fmt.Println(f * 3) // 9
	n = int(f * 3)
	fmt.Println(f, n) // 3, 9

	var n1 int32
	var n2 int8

	fmt.Println(n1, n2)
	n1 = 1024129
	n2 = int8(n1)
	fmt.Println(n1, n2)
	// 1024129 -127
	// 1024129 -> 0xfa081; 0x81
	var s string
	s = strconv.Itoa(int(n1))
	fmt.Println(s)

	n3, err := strconv.Atoi("12345")
	if err == nil {
		fmt.Println(n3)
		// 12345
	}

	var x int64
	rand.Seed(time.Now().Unix())

	x = rand.Int63()
	s = strconv.FormatInt(x, 36)
	fmt.Println(s)

	// 世界是只有10种人，一种懂二进制，不种不懂二进制。
}
