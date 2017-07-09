package main

import (
	"fmt"
	"time"
)

func main() {
	var n time.Duration
	//Duration 底层是int64类型。但用type重新定义了。

	n = time.Hour
	fmt.Println(int64(n)) // 3600000000000
	n = 3*time.Hour + 30*time.Minute
	// time.Hour, time.Minute ,n.Second ar Duration type
	// n.Minutes, n.Seconds are func
	fmt.Println("str:", n.String()) // 3h30m0s
	fmt.Println(n.Seconds)          // 0x1087b90
	fmt.Println(n.Minutes)          // 0x1087bf0

	fmt.Println(n.Seconds()) //12600
	fmt.Println(n.Minutes()) // 210. 3 * 60 +30

	t := time.Now()
	t1 := t.Add(-time.Hour)
	fmt.Println(t1.Sub(t)) // -1h0m0s
}
