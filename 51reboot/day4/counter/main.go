package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// 分割单词，统计词频
func main() {
	if len(os.Args) == 1 {
		fmt.Printf("usage: %s files", os.Args[0])
		os.Exit(1)
	}
	counter := counterWord(string(os.Args[1]))
	for word, num := range counter {
		fmt.Printf("%s\t%d\n", word, num)
	}
}

func counterWord(fileName string) map[string]int {
	buf, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	counter := make(map[string]int)
	// _ 忽略索引号，下面的if中，_是忽略value，只关注是否在map中
	for _, word := range strings.Fields(string(buf)) {
		counter[word]++
		// if _, ok := counter[word]; ok {
		// 	counter[word]++
		// } else {
		// 	counter[word] = 1
		// }
	}
	return counter
}
