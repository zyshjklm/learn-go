package main

import (
	"log"
	"os"
)

func print(s string) {
	f, err := os.Open(s)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	// 三步曲：打开，判断错误，defer关闭。
}

func main() {
	print(os.Args[1])

}
