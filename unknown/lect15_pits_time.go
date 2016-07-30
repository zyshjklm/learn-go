package main

import (
	"fmt"
	"time"
)


func main() {
	t := time.Now()
	fmt.Println(t)

	// fmt.Println(t.Format())
	fmt.Println(t.Format(time.ANSIC))
	// ANSIC = 'Mon Jan _2 15:04:05 2006'
	// refer :
	//   	 https://gowalker.org/time

	fmt.Println(t.Format("Mon Jan _2 15:04:05 2006"))

	fmt.Println(t.Format("Mon Jan _2 15:03:05 2006"))
	// 所以最好使用常量，其次是原始串。
/*
2016-07-30 11:32:53.551473043 +0800 CST
Sat Jul 30 11:32:53 2016
Sat Jul 30 11:32:53 2016
Sat Jul 30 11:11:53 2016
*/

}