// 给定参数NUM，计算从1-NUM之间所有数的fib值。
// 使用iter

package main

import (
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("usage: %s NUM(int and >=1)\n", os.Args[0])
	}
	num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal("NUM should be int and >= 1")
	}
	iter := genIter()
	for i := 1; i <= num; i++ {
		fmt.Printf("%d\tfib value: %32v\n", i, iter())
	}
}

func genIter() func() *big.Int {
	// *big.Int var for closure.
	ptmp, pcur, pnext := big.NewInt(0), big.NewInt(0), big.NewInt(1)

	return func() *big.Int {
		// printBig(&ptmp, &pcur, &pnext)
		ptmp.Set(pcur)
		pcur.Set(pnext)
		pnext.Add(pcur, ptmp)
		return pcur
	}
}

// *bit.Int is type of struct.
func printBig(x, y, z **big.Int) {
	fmt.Println("--- debug addr:", x, y, z, "val:", *x, *y, *z)
}
