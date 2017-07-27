package main

import (
	"fmt"
	"time"
)

type Account struct {
	// true allow to operation
	flag  bool
	money int
}

func (a *Account) GetGongZi(n int) {
	a.money += n
}

func (a *Account) GiveWife(n int) {
	for !a.flag {
		time.Sleep(time.Millisecond)
	}
	a.flag = false
	if a.money > n {
		a.DoPrepare()
		a.money -= n
	}
	a.flag = true
}

func (a *Account) Buy(n int) {
	for !a.flag {
		time.Sleep(time.Microsecond)
	}
	a.flag = false
	if a.money > n {
		a.DoPrepare()
		a.money -= n
	}
	a.flag = true
}
func (a *Account) Left() int {
	return a.money
}

func (a *Account) DoPrepare() {
	time.Sleep(time.Millisecond * 100)
}

func main() {
	var acc Account
	acc.flag = true

	acc.GetGongZi(10)
	go acc.GiveWife(5)
	go acc.Buy(6)

	time.Sleep(time.Millisecond * 300)
	fmt.Println(acc.Left())
	// 5
}
