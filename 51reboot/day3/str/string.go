package main

import (
	"fmt"
)

func main() {
	str1 := "hello\n\tworld\b\b"
	// "\a\b\r\n\\\\"
	// 提示声，退格，win换行，双反斜线
	doc := `
即使换行也不影响
可以验证一下
类似python的'''
`
	fmt.Println(str1)
	fmt.Println(doc)
	fmt.Println("\a")

	str1 += "gogogo"
	fmt.Println(str1)

	var c1 byte
	str2 := "golang"
	c1 = str2[0]
	fmt.Println(str2[0], c1)
	fmt.Printf("printf of c1: %d %c\n", c1, c1)
	s2 := str2[0:4]

	fmt.Println(s2)
	fmt.Println(s2[:])
	fmt.Println(str2[:])
	fmt.Println(0, len(str2)-1)

	fmt.Println("0xA")

	// ----- modify string -----
	//s2[0] = 'z'
	// error| cannot assign to s2[0]
	array := []byte(s2)
	fmt.Println(array)
	array[0] = 'H'
	fmt.Println(array)
	array[0] = 73
	fmt.Println(array)

	fmt.Println('a' + ('H' - 'h'))
	fmt.Printf("%c\n", 'a'+('H'-'h'))

	fmt.Println(toUpper("Hello"))
}

func toUpper(s string) string {
	arr := []byte(s)
	for i := 0; i < len(arr); i++ {
		if s[i] >= 97 && s[i] <= 122 {
			arr[i] -= ('a' - 'A')
		}
	}
	return string(arr)
}

func printASCII() {
	fmt.Println("\n--- outout ASCII ---")
	var b byte
	for b = 0; b < 177; b++ {
		fmt.Printf("%d %c\n", b, b)
	}
}
