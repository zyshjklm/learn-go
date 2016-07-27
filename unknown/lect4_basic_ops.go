package main 

import (
	"fmt"
)


func main() {

	fmt.Println(^2)		// 1元运算
	fmt.Println(11^2)	// 2元
	fmt.Println(1 << 10)
	fmt.Println(1 << 10 << 10 >> 8)	// 左移了12位

/*
位运算符：

 6: 0110
11: 1011
----------------
&   0010 = 2
|	1111 = 15
^	1101 = 13
&^	0100 = 4  // 第二个bit为1，则将前一个的bit置0

*/
	fmt.Println(6&11)
	fmt.Println(6|11)
	fmt.Println(6^11)
	fmt.Println(6&^11)

/*
&& 逻辑与，如果第前面的值为0，则后面的就不计算了
|| 逻辑或
*/

	for a := 5; a >=0; a-- {
		if a > 0 && (10/a) > 1 {
			fmt.Println(a, "OK")
		} else {
			fmt.Println(a, "ERR")
		}

	} 

	// 计算机存储单位的枚举 
	const (
		_ = iota
		KB float64 = 1 << (iota * 10)
		MB // Like above 
		GB 
		TB
		PB
	)

	fmt.Println()
	fmt.Println(KB)
	fmt.Println(MB)
	fmt.Println(GB)
	fmt.Println(TB)
	fmt.Println(PB)
	/*
	1024
	1.048576e+06
	1.073741824e+09
	1.099511627776e+12
	1.125899906842624e+15
	*/
}


