package main

import (
	"fmt"
	"path/filepath"
	"log"
	"net/url"
	"os"
)

func main() {
	s := os.Args[1]
	u, err := url.Parse(s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("scheme:", u.Scheme)
	// Host include Hostname and Port()
	fmt.Println("Host:", u.Host)
	fmt.Println("Hostname:", u.Hostname())
	if len(u.Port()) != 0 {
		fmt.Printf("port:%s\n", u.Port())
	}

	fmt.Println("path:", u.Path)
	fmt.Println("queryString:", u.RawQuery)
	fmt.Println("user:", u.User)
	fmt.Println("xx", u.Fragment)

	// 
	fmt.Println("dir :", filepath.Dir(u.Path))
	fmt.Println("base:", filepath.Base(u.Path))
	fmt.Println("name:", filepath.Base(filepath.Base(u.Path)))
}

/*
go run http-url.go "http://xxx.com/path/x.jpg?abcd=2342"
scheme: http
Host: xxx.com
path: /path/x.jpg
queryString: abcd=2342
user: <nil>
xx
*/
