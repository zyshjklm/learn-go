package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	cls "./class"
	stu "./student"
)

func main() {
	// test for student
	var stu1, stu2 stu.Student
	stu1.Set("jiang", 3)
	stu2.Set("xxxxx", 0)
	fmt.Println("origin struct:")
	fmt.Printf("1 %p, %v\n", &stu1, stu1)
	fmt.Printf("2 %p, %v\n", &stu2, stu2)
	buf, err := json.Marshal(stu1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\n-- ori buf: ")
	os.Stdout.Write(buf)

	json.Unmarshal(buf, &stu2)
	fmt.Printf("\n-- new stu: ")
	fmt.Println(stu2)
	fmt.Printf("1 %p, %v\n", &stu1, stu1)
	fmt.Printf("2 %p, %v\n", &stu2, stu2)
	// below is ok
	// stu2.UnmarshalJSON(buf)

	// test fo Class
	fmt.Println("\n----- start of class -----")
	var c1, c2 cls.Class
	c1.Create("zoo")
	c1.Add("jungle", 3)
	c1.Add("golang", 5)
	c1.Add("tiger", 1)
	c1.Print()

	cbuf, err := json.Marshal(c1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\n-- ori buf:\n")
	os.Stdout.Write(cbuf)

	json.Unmarshal(cbuf, &c2)
	fmt.Printf("\n-- new stu:\n")
	c2.Print()
	fmt.Println("\n----- end of class -----")
}
