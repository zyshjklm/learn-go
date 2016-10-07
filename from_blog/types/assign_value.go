package main 

import (
	"fmt"
)

type MyInt32 int32 

func main() {
	// implied convert for literal value
	var myInt int32 = 5
	var myFloat float64 = 0
	// you don't necessary to do: myFloat = float64(0)

	fmt.Println(myInt)
	fmt.Println(myFloat)

	// explicit convert for variables
	var uid int32 = 12345
	var gid int64 = int64(uid)
	// gid = uid will panic
	fmt.Printf("uid =%d, gid =%d\n", uid, gid)

	var gid2 MyInt32 = MyInt32(uid)
	fmt.Printf("uid =%d, gid2=%d\n", uid, gid2)
	// MyInt32 is static type, and int32 is the bottom type for it


	// truncation for type conversation
	// little types convert to big one is safe.
	// big to little depends on the arch of CPU
	var gid3 int32 = 0x12345678
	var uid3 int8 = int8(gid3)
	// truncation only save the lower bytes.
	// uid3 = 0x12 -> big    endian: lower bytes on higher addrs
	// uid3 = 0x78 -> little endian: lower bytes on lower addrs
	// mac is little endian.

	fmt.Printf("uid3=0x%02x, gid3=0x%02x\n", uid3, gid3)
	if uid3 == 0x12 {
		fmt.Println("big endian")
	} else  {
		fmt.Println("little endian")
	}
}


