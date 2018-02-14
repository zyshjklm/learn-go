package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	buf, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	strSlice := strings.Fields(string(buf))
	for i, j := 0, len(strSlice)-1; i < j; i, j = i+1, j-1 {
		strSlice[i], strSlice[j] = strSlice[j], strSlice[i]
	}
	fmt.Println(strings.Join(strSlice, " "))
}
