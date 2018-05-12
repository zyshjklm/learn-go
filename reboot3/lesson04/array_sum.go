package main

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	var str = "12345678"
	data := []byte(str)
	md5sum := md5.Sum(data)
	fmt.Printf("str md5:%v, len:%d\n", md5sum, len(md5sum))

	f, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	md5sum = md5.Sum(f)
	fmt.Printf("file md5:%x, len:%d\n", string(md5sum[:]), len(md5sum))
	fmt.Printf("file md5:%v, len:%d\n", md5sum, len(md5sum))
}
