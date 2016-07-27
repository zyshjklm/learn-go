package main 

import "fmt"

type test struct{}

type person struct {
	Name string
	Age int
}

func A0(per person) {
	per.Age = 13
	fmt.Println("A0", per)
}

func A1(per *person) {
	per.Age = 13
	fmt.Println("A1", per)
}

type person2 struct {
	Name string
	Age int
	Contact struct {
		Phone, City string
	}
}


func main() {
	a0 := test{}
	fmt.Println(a0)

	a1 := person{}
	fmt.Println(a1)

	a1.Name = "joe"
	a1.Age = 23
	fmt.Println(a1)

	a2 := person{ 
		Name: "kitty",
	}
	fmt.Println(a2)
	A0(a2)
	// 对a2的引用是值拷贝，不影响原结构。
	fmt.Println(a2)

	A1(&a2)
	// 对a2的引用是地址，将影响原结构。
	fmt.Println(a2)


	// struct pointer. &structName is recommend
	a3 := &person {
		Name: "Lily",
		Age: 18,
	}

	fmt.Println()
	fmt.Println(a3)
	A1(a3)
	fmt.Println(a3)

	
	// anonymous struct
	anoStr0 := struct {
		Name string
		Age int
	} {
		Name: "Gate",
		Age: 32,
	}

	fmt.Println(anoStr0)

	anoStr1 := &person2{
		Name: "Jobs",
		Age: 35,
	}
	anoStr1.Contact.Phone = "123456"
	anoStr1.Contact.City = "New-York"
	fmt.Println(anoStr1)
}


