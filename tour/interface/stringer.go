package main 

import (
	"fmt"
	"strconv"
	"strings"
)

// --------- person -----------
type Person struct {
	Name string
	Age int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}


// --------- IP -----------
type IPAddr [4]byte

func (ip IPAddr) String() string {
	var ips []string
	for _, item := range ip {
		// byte to string: byte -> int -> Itoa()
		ips = append(ips, strconv.Itoa(int(item)))
	}
	fmt.Printf("%s\n", ips)
	return fmt.Sprintf("%v", strings.Join(ips, "."))

}

func main() {
	a := Person{"Arthur Dent", 42}
	z := Person{"jungle", 33}

	fmt.Println(a, z)

	addrs := map[string]IPAddr{
		"lookback": {172, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}

	for n, ip := range addrs {
		fmt.Printf("%v: %v\n", n, ip)
	}
}