// (Package doc comment intentionally malformed to demonstrate golint.)
//package popcount
package popcount

// pc[i] is the population count of i.
var pc [256]byte
/*
note:
    1 byte is 8-bits. value from 0-255,
    so pc is 256-length.
    the max value of a pop-count is 8 of a byte.
    so type use byte.
*/

/*
    initial pc[]. 
        each byte is zero.
    for i = 0..255. pc[i] is the count if i.
    pc[0] = 0,
    for i = 1..255. pc[i] is the sum of:
        1) bit7-1 of i, -> i/2
        2) bit  0 of i, -> i%2
    so below three formula are equivalent.

    pc[i] = pc[i/2] + byte(i%2)
    pc[i] = pc[i/2] + byte(i)%2
    pc[i] = pc[i/2] + byte(i&1)
*/
func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}
/*
pc = 
[0 1 1 2 1 2 2 3 1 2 2 3 2 3 3 4 1 2 2 3 2 3 3 4 2 3 3 4 3 4 4 5 1 2 2 3 2 3 3 4 2 3 3 4 3 4 4 5 2 3 3 4 3 4 4 5 3 4 4 5 4 5 5 6 1 2 2 3 2 3 3 4 2 3 3 4 3 4 4 5 2 3 3 4 3 4 4 5 3 4 4 5 4 5 5 6 2 3 3 4 3 4 4 5 3 4 4 5 4 5 5 6 3 4 4 5 4 5 5 6 4 5 5 6 5 6 6 7 1 2 2 3 2 3 3 4 2 3 3 4 3 4 4 5 2 3 3 4 3 4 4 5 3 4 4 5 4 5 5 6 2 3 3 4 3 4 4 5 3 4 4 5 4 5 5 6 3 4 4 5 4 5 5 6 4 5 5 6 5 6 6 7 2 3 3 4 3 4 4 5 3 4 4 5 4 5 5 6 3 4 4 5 4 5 5 6 4 5 5 6 5 6 6 7 3 4 4 5 4 5 5 6 4 5 5 6 5 6 6 7 4 5 5 6 5 6 6 7 5 6 6 7 6 7 7 8]
*/

// PopCount returns the population count (number of set bits) of x.
// split x to 8 byte
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

