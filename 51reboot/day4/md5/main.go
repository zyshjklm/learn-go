package main

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("usage: %s file1 file2 ...", os.Args[0])
		os.Exit(1)
	}
	//下标从0到长度-1，故用< len()
	for i := 1; i < len(os.Args); i++ {
		// 此处用ReadFile，只用做示例，线上程序不能这样用，有耗尽内存的风险
		buf, err := ioutil.ReadFile(os.Args[i])
		if err != nil {
			fmt.Println("read error of", os.Args[i])
		}
		md5sum := md5.Sum(buf)
		fmt.Printf("%x %s\n", md5sum, os.Args[i])
	}
}
