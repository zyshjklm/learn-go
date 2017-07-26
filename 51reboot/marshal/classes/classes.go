package classes

import (
	"encoding/json"
	"fmt"
	"reflect"

	cls "../class"
)

type Classes struct {
	allClasses   map[string]*cls.Class `json:"classes"`
	curClassName string                `json:"current`
}

// Class struct for a group of student in the same Class
// type Class struct {
// 	allStudents []*stu.Student `json:"students`
// 	className   string         `json:"name"`
// }

// Create create a class struct.
func (clses *Classes) Create(className string) error {
	if len(clses.allClasses) == 0 {
		clses.allClasses = make(map[string]*cls.Class)
	}
	for curName := range clses.allClasses {
		if curName == className {
			return fmt.Errorf("duplicated class name")
		}
	}
	var cls cls.Class
	cls.Create(className)
	clses.allClasses[className] = &cls
	clses.curClassName = className
	return nil
}

// Add add a new student to current class
func (clses *Classes) Add(name string, id int) error {
	if len(name) == 0 {
		return fmt.Errorf("name is blank")
	}
	// var tmp stu.Student
	// tmp.Set(name, id)
	clses.allClasses[clses.curClassName].Add(name, id)
	// curClass := clses.allClasses[clses.curClassName]
	// curClass = append(curClass, &tmp)
	return nil
}

// Print print all stu
func (clses *Classes) Print() {
	fmt.Printf("-- classes info:\n")
	for curName := range clses.allClasses {
		fmt.Printf("class %s:\n", curName)
		clses.allClasses[curName].Print()
		// fmt.Printf("%d\t%v\n", i, cls.allStudents[i])
	}
}

// MarshalJSON my own marshal.
// cls must be a struct copy. not pointer
func (clses Classes) MarshalJSON() ([]byte, error) {
	fmt.Println("clses ma-json:", reflect.TypeOf(clses))
	// allClasses   map[string]*cls.Class `json:"classes"`
	// curClassName string                `json:"current`
	tmpStruct := make(map[string]interface{})
	tmpStruct["AllClasses"] = clses.allClasses
	tmpStruct["CurClassName"] = clses.curClassName
	return json.Marshal(tmpStruct)
	// 上面的办法看上去更优雅。但即不方便定义序列化时显示的key。
	// 下面的方法，可以统一定义一个临时结构，供MarshalJSON和UnmarshalJSON，
	// 这样就不会显得太冗余。

	return json.Marshal(struct {
		AllClasses   map[string]*cls.Class `json:"classes`
		CurClassName string                `json:"current"`
	}{
		AllClasses:   clses.allClasses,
		CurClassName: clses.curClassName,
	})
}

// UnmarshalJSON un
// cls must be a pointer of Student struct
func (clses *Classes) UnmarshalJSON(data []byte) error {
	var tmp struct {
		AllClasses   map[string]*cls.Class `json:"classes`
		CurClassName string                `json:"current"`
	}
	if err := json.Unmarshal(data, &tmp); err != nil {
		fmt.Println("err of UN,", err)
		return err
	}
	clses.allClasses = tmp.AllClasses
	clses.curClassName = tmp.CurClassName
	return nil
}
