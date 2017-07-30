package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func getFD() *os.File {
	f, err := os.Open("a.txt")
	if err != nil {
		log.Fatal(err)
	}
	return f
}
func main() {
	var offSet, num int64

	fd1 := getFD()
	defer fd1.Close()

	offSet, _ = fd1.Seek(0, 1)
	log.Print("current offset:", offSet)

	r1 := bufio.NewReaderSize(fd1, 10)
	line1, _ := r1.ReadString('\n')
	offSet, _ = fd1.Seek(0, 1)
	log.Print("current offset:", offSet)

	fmt.Printf("#%s#\n", line1)
	// 从效果上讲，先读了第一行到行末的换行符；然后又读了10个字符
	// 因此此时只剩下2个字符"ng"没有读了。
	num, _ = io.Copy(os.Stdout, fd1)
	offSet, _ = fd1.Seek(0, 1)
	log.Printf("current offset:%d; num:%d\n", offSet, num)

	log.Print("\n--- test fd2 ---\n")

	fd2 := getFD()
	defer fd2.Close()
	offSet, _ = fd2.Seek(0, 1)
	log.Print("current offset:", offSet)

	r2 := bufio.NewReaderSize(fd2, 10)
	line2, _ := r2.ReadString('\n')
	offSet, _ = fd2.Seek(0, 1)
	log.Print("current offset:", offSet)

	fmt.Printf("#%s#\n", line2)
	num, _ = io.Copy(os.Stdout, r2)
	offSet, _ = fd2.Seek(0, 1)
	log.Printf("current offset:%d; num:%d\n", offSet, num)
}

// 'ng' vs 'hello golang'
