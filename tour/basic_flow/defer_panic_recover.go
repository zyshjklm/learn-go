// demonstrates the mechanics of panic and defer.
package main 

import (
	"fmt"
)

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking when i=", i)
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("-- Defer in g :", i)
	fmt.Println("Printing in g :", i)
	g(i+1)
}

func f() {
	// if remove the deferred func, the panic is not recovered
	// and reach the top of the goroutine's call stack.
	// terminating the program. 
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f with panic value :", r)
		}
	}()

	fmt.Println("Calling g(0).")
	g(0)
	fmt.Println("Returned normally from g(0).")
}

func main() {
	fmt.Println("Calling f().")
	f()
	fmt.Println("Returned normally from f()")
}

// no recover result :
/*
panic: 4

goroutine 1 [running]:
panic(0xb8c20, 0xc82006a2e0)
	/usr/local/go/src/runtime/panic.go:481 +0x3e6
main.g(0x4)
	/Users/user/_go/src/github.com/jungle85gopy/learn-go/tour/basic_flow/defer_panic_recover.go:11 +0x28b
main.g(0x3)
	/Users/user/_go/src/github.com/jungle85gopy/learn-go/tour/basic_flow/defer_panic_recover.go:15 +0x583
main.g(0x2)
	/Users/user/_go/src/github.com/jungle85gopy/learn-go/tour/basic_flow/defer_panic_recover.go:15 +0x583
main.g(0x1)
	/Users/user/_go/src/github.com/jungle85gopy/learn-go/tour/basic_flow/defer_panic_recover.go:15 +0x583
main.g(0x0)
	/Users/user/_go/src/github.com/jungle85gopy/learn-go/tour/basic_flow/defer_panic_recover.go:15 +0x583
main.f()
	/Users/user/_go/src/github.com/jungle85gopy/learn-go/tour/basic_flow/defer_panic_recover.go:29 +0xeb
main.main()
	/Users/user/_go/src/github.com/jungle85gopy/learn-go/tour/basic_flow/defer_panic_recover.go:35 +0xe3
exit status 2

*/
