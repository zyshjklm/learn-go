// practise 3.5.2 unicode
package main 

import (
	"fmt"
	"unicode/utf8"
)


func HasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

func HasSuffix(s, suffix string) bool {
	return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
}

func Contains(s, substr string) bool {
	for i := 0; i <= len(s) - len(substr); i++ {
		if HasPrefix(s[i:], substr) {
			return true
		}
	}
	return false
}

func main() {


	s1 := "hello world"
	s2 := "world"
	fmt.Println(s1, s2)
	fmt.Println(Contains(s1, s2))
	
	// unicode
	unicode := string('世')
	fmt.Println(unicode[:])
	
	str_len := len(unicode)
	uni_len := utf8.RuneCountInString(unicode)

	fmt.Println(str_len)
	fmt.Println(uni_len)
	for i := 0; i < str_len; i++ {
		fmt.Printf("%08b\n", unicode[i])
	}
	// 11100100 10111000 10010110

	s := "hello 世界!"

	fmt.Println("traversal str :", s)
	for i:= 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c\t%d\n", i, r, size)
		i += size
	}
	
	fmt.Println("range str :", s)
	for i, r := range s {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}

}

