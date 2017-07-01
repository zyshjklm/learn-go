package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

// 错误处理的三种方法：
// 1）直接退出，如程序启动时打开进程文件，log.Fatal(err)
// 2) 重试。如网络超时。
// 3）直接返回err，将异常上抛。

func main() {
	f, err := os.Open("a.txt")
	if err != nil {
		log.Fatal(err) // log and exit
	}

	var cont string
	retries := 3
	for i := 1; i <= retries; i++ {
		cont, err = read(f)
		if err == nil {
			break
		}
		time.Sleep(time.Second << uint64(i))
	}
	fmt.Println(cont)
}

func read(f *os.File) (string, error) {
	var total []byte
	buf := make([]byte, 1024)
	for {
		n, err := f.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			return "", err
		}
		total = append(total, total[:n]...)
	}
	return string(total), nil
}
