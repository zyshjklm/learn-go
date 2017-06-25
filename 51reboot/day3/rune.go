package main

import "fmt"

func main() {
	// byte = uint8
	// rune = uint32

	s := "golang你好"
	fmt.Println(len(s))
	// 12 bytes. len return byte number.

	cnt := 0
	// range 以字符为单位，长度为8
	for _, r := range s {
		fmt.Printf("%c\n", r)
		cnt++
	}
	fmt.Println("cnt:", cnt)

	cnt = 0
	// range 以byte为单位，长度为12
	for _, r := range []byte(s) {
		fmt.Printf("%c\n", r)
		cnt++
	}
	fmt.Println("cnt:", cnt)

	ss := []rune("hello")
	cnt = 0
	// rune为单位，长度为5
	for _, r := range ss {
		fmt.Printf("%c\n", r)
		cnt++
	}
	fmt.Println("cnt:", cnt)

	ss = []rune("hello中国")
	cnt = 0
	// rune为单位，长度为7
	for _, r := range ss {
		fmt.Printf("%c\n", r)
		cnt++
	}
	fmt.Println("cnt:", cnt)
}
