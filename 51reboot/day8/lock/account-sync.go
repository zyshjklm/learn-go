package main

import (
	"fmt"
	"sync"
	"time"
)

type Account struct {
	flag  sync.Mutex
	money int
}

func (a *Account) GetGongZi(n int) {
	a.money += n
}

func (a *Account) GiveWife(n int) {
	a.flag.Lock()
	if a.money > n {
		a.DoPrepare()
		a.money -= n
	}
	a.flag.Unlock()
}

func (a *Account) Buy(n int) {
	a.flag.Lock()
	if a.money > n {
		a.DoPrepare()
		a.money -= n
	}
	a.flag.Unlock()
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
	go acc.Buy(5)
	go acc.GiveWife(6)

	time.Sleep(time.Millisecond * 300)
	fmt.Println(acc.Left())
	// 5
}
