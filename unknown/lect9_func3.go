package main 

import "fmt"

func main() {
	// test defer in for
	// defer是基于函数级别，每个匿名函数结果时即执行defer
	// 而不用等到整个for结束

	for i := 0; i < 5; i++ {
		func() {
			fmt.Println("func, i =", i)
			// run after current anonymous func finish.
			defer fmt.Println(" -- func DEFER, i =", i)
		}()
		fmt.Println()
		// run after main finish
		defer fmt.Println(" ====== after func DEFER, i =", i)
	}

}

/*
result as below:

func, i = 0
 -- func DEFER, i = 0

func, i = 1
 -- func DEFER, i = 1

func, i = 2
 -- func DEFER, i = 2

func, i = 3
 -- func DEFER, i = 3

func, i = 4
 -- func DEFER, i = 4

 ====== after func DEFER, i = 4
 ====== after func DEFER, i = 3
 ====== after func DEFER, i = 2
 ====== after func DEFER, i = 1
 ====== after func DEFER, i = 0

 */
 
