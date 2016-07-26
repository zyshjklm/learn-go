package main 

import "fmt"

type A struct {
	Name string
}

type B struct {
	Name string
}

func (a A) Print(x int) {
	fmt.Println("A", x)
}

func (b B) Print() {
	fmt.Println("B")
}


// pointer
func (a *A) PrintName() {
	a.Name = "AAA"
	fmt.Println("A")
}

func (b B) PrintName() {
	b.Name = "BBB"
	fmt.Println("B")
}

func main() {
	a := A{}
	a.Print(3)

	b := B{}
	b.Print()

	a1 := A{}
	a1.PrintName()
	fmt.Println("a1.Name:", a1.Name)

	b1 := B{}
	b1.PrintName()
	fmt.Println("b1.Name:", b1.Name)
}