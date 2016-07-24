package main 

import "fmt"

// slice

func main() {
	var s0 []int	// define slice
	fmt.Println(s0)	// []

	a := [10]int{0,1,2,3,4,5,6,7,8,9}
	fmt.Println(a)

	s1 := a[6]		// value of 6
	fmt.Println(s1)

	s2 := a[6:7]	// slice [6]
	fmt.Println(s2)

	s3 := a[5:]
	fmt.Println(s3)

	s4 := a[5:len(a)]
	fmt.Println(s4)

	s5 := a[:5]
	fmt.Println(s5, "\n")

	// make 
	sm0 := make([]int, 3, 10)	// type, len, cap
	fmt.Println(sm0)
	fmt.Println(len(sm0))
	fmt.Println(cap(sm0))

	fmt.Println()
	sm1 := []byte{'a', 'b','c','d','e','f','g','h','i','j','k'}
	sl0 := sm1[2:5]
	fmt.Println("byte:", sl0)
	fmt.Println("str :", string(sl0))
	fmt.Println("len of sm1:", len(sm1))
	fmt.Println("len of sl0:", len(sl0))
	fmt.Println("cap of sl0:", cap(sl0))
	// sl0指向了byte数组的c-k, 故有len和cap

	// 同上，因此sl0可以超过其本来的3个长度取值
	sl1 := sl0[4:7]
	fmt.Println(sl1)

	// reslice
	sl2 := sl0[1:3]
	sl3 := sl0[3:5]
	fmt.Println(string(sl2))
	fmt.Println(string(sl3))
	// 索引时不要超过cap的长度。
	// 否则会引起超界，并引发错误

}
