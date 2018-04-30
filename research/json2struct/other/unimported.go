package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// Student struct for student info. id is unique.
type Student struct {
	stuID   int    `json:"id"`
	stuName string `json:"name"`
	stuAge  int    `json:"age"`
}

// MarshalJSON my own marshal.
// stu must be a struct copy. not pointer
func (stu Student) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"id":   stu.stuID,
		"name": stu.stuName,
		"age":  stu.stuAge,
	})
}

// UnmarshalJSON un
// stu must be a pointer of Student struct
func (stu *Student) UnmarshalJSON(data []byte) error {
	var tmpStu struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	if err := json.Unmarshal(data, &tmpStu); err != nil {
		fmt.Println("err:", err)
		return err
	}
	stu.stuID = tmpStu.ID
	stu.stuName = tmpStu.Name
	stu.stuAge = tmpStu.Age
	return nil
}

func main() {
	var stu = Student{}
	stu.stuName = "jungle85"
	stu.stuID = 1
	stu.stuAge = 23

	buf, _ := json.Marshal(stu)
	log.Println(string(buf))

	stuOut := new(Student)
	if err := json.Unmarshal(buf, stuOut); err != nil {
		log.Fatal(err)
	}
	log.Printf("%#v\n", stuOut)
}
