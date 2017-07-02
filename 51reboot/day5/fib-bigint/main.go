// 给定参数NUM，计算从1-NUM之间所有数的fib值。
// 使用map保存中间结果，加速计算

package main

import (
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"
)

// refer: http://localhost:6060/pkg/math/big/#Int.Bit
var fibs map[int]*big.Int

func init() {
	fibs = make(map[int]*big.Int, 1024)
	fibs[1] = big.NewInt(1)
	fibs[2] = big.NewInt(1)
}

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("usage: %s NUM(int and >=1)\n", os.Args[0])
	}
	num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal("NUM should be int and >= 1")
	}
	fmt.Println(fibs)
	searchFib(num)
}

func searchFib(end int) {
	for i := 1; i <= end; i++ {
		if _, ok := fibs[i]; !ok {
			// fmt.Println("call calcfib() with idx:", i)
			calcFib(i)
		}
		fmt.Printf("%d\tfib value: % 32v\n", i, fibs[i])
	}
}

func calcFib(idx int) {
	if idx <= 2 {
		log.Fatal("index too little!")
	}
	if _, ok := fibs[idx-1]; !ok {
		calcFib(idx - 1)
	}
	if _, ok := fibs[idx-2]; !ok {
		calcFib(idx - 2)
	}
	// fmt.Printf("%v, %v\n", fibs[idx-1], fibs[idx-2])

	fibs[idx] = big.NewInt(0)
	// default value of fibs[idx] is nil. so need to be inited by zero. otherwise:
	// panic: runtime error: invalid memory address or nil pointer dereference
	// [signal SIGSEGV: segmentation violation code=0x1 addr=0x18 pc=0x1094078]
	fibs[idx].Add(fibs[idx-1], fibs[idx-2])
}
