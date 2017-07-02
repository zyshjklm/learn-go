// 给定参数NUM，计算从1-NUM之间所有数的fib值。
// 使用map保存中间结果，加速计算
// 当NUM小于93时工作正常。之后的计算结果将溢出。
package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

var fibs map[int]uint64

func init() {
	fibs = make(map[int]uint64, 1024)
	fibs[1] = 1
	fibs[2] = 1
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
		// fmt.Println("idx: ", i)
		if _, ok := fibs[i]; !ok {
			calcFib(i)
		}
		fmt.Printf("%d\tfib value: % 24d\n", i, fibs[i])
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
	fibs[idx] = fibs[idx-1] + fibs[idx-2]
}
