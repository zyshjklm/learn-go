package main

import "log"

func main() {
	cnt := 0
	for {
		var n int32
		go func() {
			n = n + 2
		}()
		go func() {
			n = n / 2
		}()
		cnt++
		if n == 2 {
			log.Fatalf("counter:%d, %d\n", cnt, n)
			// panic("bingo")
		}
	}
}
