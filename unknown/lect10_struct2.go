package main 

import "fmt"


type human struct {
	Sex byte
}

// struct 组合
type teacher struct {
	human	// anonymous struct embed
	Name string
	Age int
}

type student struct {
	human
	Name string
	Age int
}


type A struct {
	B 
	C
	Name string
}

type B struct {
	Name string
	Age int
	Height float32
}
type C struct {
	Name string
	Age int
}

func main() {

	ta := teacher{Name: "Joe", Age: 29,
		human: human{Sex: 0},
	}
	sa := student{Name: "bob", Age: 18,
		human: human{Sex: 1},
	}

	fmt.Println(ta, sa)

	ta.human.Sex = 9
	sa.Sex = 8
	fmt.Println(ta, sa)

	
	// same filed name in both 
	// anonymous filed and outside struct
	a := A{Name: "A", 
		B: B{Name: "B", Age: 22, Height: 1.73},
		C: C{Name: "C", Age: 11},
	}
	fmt.Println(a.Name, a.B.Name, a.C.Name)
	fmt.Println(a.Height)
	fmt.Println(a.B.Age, a.C.Age)
}


