package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	inStr string
)

const (
	colorGreen = "\033[32;1m%v\033[0m\n"
	colorRed   = "\033[31;1m%v\033[0m\n"
)

func main() {
	f := bufio.NewReader(os.Stdin)
	num := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("number："+colorGreen, num)
	for {
		fmt.Print("input a number as start position>")
		line, _ := f.ReadString('\n')
		if len(line) == 1 {
			continue
		}
		fmt.Sscan(line, &inStr)
		if inStr == "stop" {
			break
		}
		index, err := strconv.Atoi(inStr)
		if err != nil {
			fmt.Println("\t[error] should be a number")
			continue
		}
		if index < 0 || index >= len(num) {
			fmt.Println("\t[error] index is out of range!")
			continue
		}
		numStd := num[:index]
		numEnd := num[index:]
		for i := 0; i < len(numStd); i = i + 1 {
			numEnd = append(numEnd, numStd[i])
		}
		fmt.Printf("number："+colorRed, numEnd)
	}
}
