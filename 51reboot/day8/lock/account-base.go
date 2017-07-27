package main

import (
	"fmt"
	"time"
)

type Account struct {
	money int
}

func (a *Account) GetGongZi(n int) {
	a.money += n
}

func (a *Account) GiveWife(n int) {
	if a.money > n {
		a.DoPrepare()
		a.money -= n
	}
}

func (a *Account) Buy(n int) {
	if a.money > n {
		a.DoPrepare()
		a.money -= n
	}
}
func (a *Account) Left() int {
	return a.money
}

func (a *Account) DoPrepare() {
	time.Sleep(time.Millisecond * 100)
}

func main() {
	var acc Account
	acc.GetGongZi(10)
	go acc.GiveWife(5)
	go acc.Buy(6)

	time.Sleep(240 * time.Millisecond)
	fmt.Println(acc.Left())
	// -1
}
