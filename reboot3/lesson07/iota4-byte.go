package main

import "fmt"

type ByteSize float64

const (
	_           = iota
	KB ByteSize = 1 << (iota * 10)
	MB
	GB
	TB
)

func main() {
	fmt.Println(KB, MB, GB)
}

/*
go run iota4-byte.go
1024 1.048576e+06 1.073741824e+09

*/
