// refer from: gopl/ch2/popcount/popcount.go
//package popcount
package popcount

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(src [32]uint8) int {
    var sum int
    for i := range src {
        sum += int(pc[byte(i)])
    }
	return sum
}

func PopCount48(src [48]uint8) int {
    var sum int
    for i := range src {
        sum += int(pc[byte(i)])
    }
	return sum
}

func PopCount64(src [64]uint8) int {
    var sum int
    for i := range src {
        sum += int(pc[byte(i)])
    }
	return sum
}

func PopCountSlice(src []uint8) int {
    var sum int
    for i := range src {
        sum += int(pc[byte(i)])
    }
	return sum
}
