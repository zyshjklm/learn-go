package main 

import (
	"fmt"
	"crypto/sha256"
	"crypto/sha512"
    "github.com/jungle85gopy/learn-go/gopl/ch4/popcount"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))

	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
	// output:
	// 2d711642b726b04401627ca9fbac32f5c8530fb1903cc4db02258717921a4881
	// 4b68ab3847feda7d6c62c1fbcbeebfa35eab7351ed5e78f4ddadea5df64b8015
	// false
	// [32]uint8

// exercise 4.1
    fmt.Printf("length: %d of c1\n", popcount.PopCount(c1))
    fmt.Printf("length: %d of c2\n", popcount.PopCount(c2))
    fmt.Printf("length: %d of c1\n", popcount.PopCountSlice(c1[:]))
    fmt.Printf("length: %d of c2\n", popcount.PopCountSlice(c2[:]))

// exercise 4.2 
	c3 := sha512.Sum384([]byte("x"))
	c4 := sha512.Sum512([]byte("X"))

	fmt.Printf("%x\n%x\n%T\n%T\n", c3, c4, c3, c4)
	//
	// d752c2c51fba0e29aa190570a9d4253e44077a058d3297fa3a5630d5bd012622f97c28acaed313b5c83bb990caa7da85
	// 3173f0564ab9462b0978a765c1283f96f05ac9e9f8361ee1006dc905c153d85bf0e4c45622e5e990abcf48fb5192ad34722e8d6a723278b39fef9e4f9fc62378
	// [48]uint8
	// [64]uint8
    fmt.Printf("length: %d of c3\n", popcount.PopCount48(c3))
    fmt.Printf("length: %d of c4\n", popcount.PopCount64(c4))
    fmt.Printf("length: %d of c3\n", popcount.PopCountSlice(c3[:]))
    fmt.Printf("length: %d of c4\n", popcount.PopCountSlice(c4[:]))
}
