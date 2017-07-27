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
	fin := make(chan int)

	acc.GetGongZi(10)
	go func() {
		acc.GiveWife(6)
		fin <- 6
	}()
	go func() {
		acc.Buy(5)
		fin <- 5
	}()

	<-fin
	<-fin
	fmt.Println(acc.Left())

}
