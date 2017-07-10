package student

import (
	"encoding/json"
	"fmt"
)

// Student struct for student info. id is unique.
type Student struct {
	stuId   int    `json:"id"`
	stuName string `json:"name"`
}

// Set init a student info
func (stu *Student) Set(name string, id int) error {
	stu.stuId = id
	stu.stuName = name
	return nil
}

// MarshalJSON my own marshal.
// stu must be a struct copy. not pointer
func (stu Student) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Age  int    `json:"age"`
	}{
		ID:   stu.stuId,
		Name: stu.stuName,
		Age:  12,
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
	stu.stuId = tmpStu.ID
	stu.stuName = tmpStu.Name
	return nil
}
