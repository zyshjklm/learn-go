package main

import "fmt"

// Student for stu
type Student struct {
	Name string
	ID   int
}

// Update for update stu id
func (s *Student) Update(id int) {
	s.ID = id
}

func main() {
	var f func(int)
	// f = Student.Update
	// error usage as above. ok as below
	// not use struct, use instance of struct.

	s := Student{Name: "binggan", ID: 3}
	f = s.Update
	f(9)
	fmt.Println(s)
	// {binggan 9}

	// 面向对象，延迟绑定
	var f1 func(s *Student, id int)
	f1 = (*Student).Update
	f1(&s, 5)
	fmt.Println(s)
	// {binggan 5}

	// 上面传的s，这里传的s1，用于将f1方法绑定到不同的对象上。
	s1 := Student{Name: "jack"}
	f1(&s1, 4)
	fmt.Println(s1)
	// {jack 4}
}
