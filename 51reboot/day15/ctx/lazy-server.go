package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

var (
	quick = "quick response"
	slow  = "slow response"
)

func lazyHandler(w http.ResponseWriter, req *http.Request) {
	ranNum := rand.Intn(2)
	if ranNum == 0 {
		time.Sleep(4 * time.Second)
		fmt.Fprintf(w, "%s, %d\n", slow, ranNum)
		fmt.Printf("%s, %d\n", slow, ranNum)
		return
	}
	fmt.Fprintf(w, "%s, %d\n", quick, ranNum)
	fmt.Printf("%s, %d\n", quick, ranNum)
	return
}

func main() {
	http.HandleFunc("/", lazyHandler)
	http.ListenAndServe(":8021", nil)
}
