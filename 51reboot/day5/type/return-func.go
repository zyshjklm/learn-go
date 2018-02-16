package main

import "fmt"

func print1() {
	fmt.Println("print1")
}
func print2() {
	fmt.Println("print2")
}

func main() {
	var f func()
	var flist [3]func()
	// var fslice []func()

	f = print1
	f()
	f = print2
	f()

	flist[0], flist[1], flist[2] = nil, print1, print2
	for i := 0; i < len(flist); i++ {
		if flist[i] != nil {
			flist[i]()
		}
	}

	fmap := map[string]func(){
		"print1": print1,
		"print2": print2,
	}
	fmap["print1"]()

}
