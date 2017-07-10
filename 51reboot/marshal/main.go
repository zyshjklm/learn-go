package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	stu "./student"
)

func main() {
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
}
