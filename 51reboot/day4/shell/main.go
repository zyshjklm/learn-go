package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	doLoop1()
	doLoop2()
}

func doLoop1() {
	var str string
	var n int
	for {
		fmt.Print("> ")
		fmt.Scan(&str, &n)
		if str == "stop" {
			break
		}
		fmt.Println(str, n)
	}
}
func doLoop2() {
	var line string
	f := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		line, _ = f.ReadString('\n')
		line = strings.Replace(line, "\n", "", 1)
		if line == "stop" {
			break
		}
		fmt.Println(line, "len:", len(line))
	}
}
