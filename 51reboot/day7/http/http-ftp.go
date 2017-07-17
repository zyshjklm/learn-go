package main

import (
	"log"
	"net/http"
	"os"
)

// 使用当前目录作为文件服务器
func main() {
	if len(os.Args) != 2 {
		log.Fatalf("usage: %s ':port'\n", os.Args[0])
	}
	http.Handle("/", http.FileServer(http.Dir(".")))
	log.Fatal(http.ListenAndServe(os.Args[1], nil))
}

/*
go run http-ftp.go ":9090"

curl localhost:9090

*/
