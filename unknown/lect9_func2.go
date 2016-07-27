package main 

import "fmt"

func closureOut(x int) func(int) int {
	fmt.Printf("%p\n", &x)
	return func(y int) int {
		fmt.Printf("%p\n", &x)
		return x + y
	}
}

func defer1() {
 	// defer
 	fmt.Println()
	fmt.Println("a")
	defer fmt.Println("b")
	fmt.Println("c")
	// a c b 
}

func defer2() {
 	fmt.Println()
	for i := 0; i < 3; i++ {
		defer fmt.Println(i)
	}

}

func defer3() {
 	fmt.Println()
	for i := 0; i < 3; i++ {
		defer func() {
			fmt.Println(i)
		}()
		// anonymous ref i
		// defer and anonymous. output: 3 3 3 
	}

}

func panicBefore() {
	fmt.Println("before panic...")
}

// error handle
func panicing() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recover in panicing...")
		}
	}()

	panic("Panic in panicing().")
}

func panicAfter() {
	fmt.Println("after panic...")
}

// exercise
func test() {
	var funcArr = [4]func(){}

	for i := 0; i < 4; i++ {
		defer fmt.Println("defer i =", i)
		defer func() {
			fmt.Println("defer_closure i =", i)
		}()
		funcArr[i] = func() {
			fmt.Println("closure i =", i)
			// i ref from var i in for.
		}
	}

	for _, f := range funcArr {
		f()
	}
/*  result:
	closure i = 4
	closure i = 4
	closure i = 4
	closure i = 4
	defer_closure i = 4
	defer i = 3
	defer_closure i = 4
	defer i = 2
	defer_closure i = 4
	defer i = 1
	defer_closure i = 4
	defer i = 0
*/

}


func main() {
	// 匿名函数
	a := func() {
		fmt.Println("func with anonymous...")
	}
	a()

	// closure
	closureIn := closureOut(10)
	fmt.Println(closureIn(1))
	fmt.Println(closureIn(3))

	fmt.Println("\n-- in one --")
	fmt.Println(closureOut(3)(4))

	defer1()
	defer2()
	defer3()

	panicBefore()
	panicing()
	panicAfter()

	fmt.Println("\n--- test result: ---")
	test()
}


