package main

import (
	"fmt"
	"strings"
)

// BitFlag for big flags
type BitFlag int

const (
	// Active use bit 0
	Active BitFlag = 1 << iota
	// Send use bit 1
	Send
	// Receive use bit 2
	Receive
)

var flag = Send | Active

func (flag BitFlag) String() string {
	var flags []string
	if flag&Active == Active {
		flags = append(flags, "Active")
	}
	if flag&Send == Send {
		flags = append(flags, "Send")
	}
	if flag&Receive == Receive {
		flags = append(flags, "Receive")
	}
	if len(flags) > 0 {
		return fmt.Sprintf("%d(%s)", int(flag), strings.Join(flags, "|"))
	}
	return "0()"
}

func main() {
	for i := 0; i < 16; i++ {
		fmt.Println(i, ":", BitFlag(i))
	}
}
