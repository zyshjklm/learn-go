package class

import (
	"encoding/json"
	"fmt"
	"reflect"

	stu "../student"
)

// Class struct for a group of student in the same Class
type Class struct {
	allStudents []*stu.Student `json:"students`
	className   string         `json:"name"`
}

// Create create a class struct. set name and
func (cls *Class) Create(s string) {
	cls.className = s
}

// Add add a new student
func (cls *Class) Add(name string, id int) error {
	if len(name) == 0 {
		return fmt.Errorf("name is blank")
	}
	var tmp stu.Student
	tmp.Set(name, id)
	cls.allStudents = append(cls.allStudents, &tmp)
	return nil
}

// Print print all stu
func (cls *Class) Print() {
	fmt.Printf("-- class [%s] students info:\n", cls.className)
	for i := 0; i < len(cls.allStudents); i++ {
		fmt.Printf("%d\t%v\n", i, cls.allStudents[i])
	}
}

// MarshalJSON my own marshal.
// cls must be a struct copy. not pointer
func (cls Class) MarshalJSON() ([]byte, error) {
	fmt.Println("cls ma-json:", reflect.TypeOf(cls))
	return json.Marshal(struct {
		AllStudents []*stu.Student
		ClassName   string `json:"name"`
	}{
		AllStudents: cls.allStudents,
		ClassName:   cls.className,
	})
}

// UnmarshalJSON un
// cls must be a pointer of Student struct
func (cls *Class) UnmarshalJSON(data []byte) error {
	var tmp struct {
		AllStudents []*stu.Student
		ClassName   string `json:"name"`
	}
	if err := json.Unmarshal(data, &tmp); err != nil {
		fmt.Println("err of UN,", err)
		return err
	}
	cls.allStudents = tmp.AllStudents
	cls.className = tmp.ClassName
	return nil
}
