package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	read4()
	read5()
}

func read4() {
	f, err := os.Open("a.txt")
	if err != nil {
		log.Fatal(err)
	}
	// 相对于当文件开始位置，偏移10个字节
	f.Seek(10, os.SEEK_SET)
	buf := make([]byte, 64)
	f.Read(buf)
	fmt.Println(buf)
}

func read5() {
	f, err := os.Open("a.txt")
	if err != nil {
		log.Fatal(err)
	}
	buf := make([]byte, 64)
	f.Read(buf)
	fmt.Println(buf)

	f.Seek(0, os.SEEK_SET)
	r := bufio.NewReader(f)
	for {
		// 按行读取
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Print(line)
	}
	f.Close()
}
