package main

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/cpu"
)

func main() {
	fmt.Println(time.Now().String())
	cores, err := cpu.Percent(time.Second*5, false)
	fmt.Println(time.Now().String())
	if err != nil {
		panic(err)
	}
	fmt.Println(cores)

	fmt.Println(time.Now().String())
	cores, err = cpu.Percent(time.Second*5, true)
	fmt.Println(time.Now().String())
	if err != nil {
		panic(err)
	}
	fmt.Println(cores)
	fmt.Println(time.Now().String())
}
