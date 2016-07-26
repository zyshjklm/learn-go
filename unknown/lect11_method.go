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
	fmt.Println(a.Name)
}

func (b B) PrintName() {
	b.Name = "BBB"
	fmt.Println("B")
	fmt.Println(b.Name)
}

type TINT int 

func (tint *TINT) Print() {
	fmt.Println("TINT...")
}
// 通过type来定义自己的Int, 并实现相应的method。

func (tint *TINT) Increase() {
	*tint += (TINT)(100)
	fmt.Println(*tint)
}

type Mod struct {
	Name string
}

// method access type field
func (m *Mod) ModifyField() {
	m.Name = "modify-by-method"
	fmt.Println(m.Name)
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
	// PrintName not modify b1.Name

	var t TINT
	t.Print()
	(*TINT).Print(&t)

	// modify field
	m := Mod{}
	m.ModifyField()
	m.Name = "modify-by-main"
	fmt.Println(m.Name)

	var t1 TINT
	t1 = 3
	fmt.Println(t1)
	t1.Increase()
	fmt.Println(t1)
}

