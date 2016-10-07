package main 

import (
	"fmt"
)

type Speaker interface {
	Say(string)
	Listen(string) string
	Interrupt(string)
}

type WangLan struct {
	msg string
}

func (this *WangLan) Say(msg string) {
	fmt.Printf("WangLan: %s\n", msg)
}

func (this *WangLan) Listen(msg string) string {
	this.msg = msg
	return msg
}

func (this *WangLan) Interrupt(msg string) {
	this.Say(msg)
}

type JiangLou struct {
	msg string
}

func (this *JiangLou) Say(msg string) {
	fmt.Printf("JiangLou: %s\n", msg)
}

func (this *JiangLou) Listen(msg string) string {
	this.msg = msg
	return msg
}

func (this *JiangLou) Interrupt(msg string) {
	this.Say(msg)
}

func main() {
	// common types convert to interface type 
	var val interface{} = "hello"
	fmt.Println(val)

	// implit convert 
	val = []byte{'a', 'b', 'c'}
	fmt.Println(val)

	// interface convert 
	wl := &WangLan{}
	jl := &JiangLou{}

	// interface variable
	var person Speaker
	person = wl 
	person.Say("Hello World!")

	person = jl
	person.Say("Good Luck!")
}


