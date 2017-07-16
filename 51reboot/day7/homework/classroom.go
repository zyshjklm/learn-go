package main

import (
	"encoding/json"
	"fmt"
	"os"
)

var classrooms map[string]*ClassRoom

type Student struct {
	name string
	id   int
}

// ----
type ClassRoom struct {
	students     map[string]*Student
	currentClass string
}

// 对于简单形式的数据，直接用其核心部分进行序列化
// 对于复杂的数据，先定义一个变量，其内部包括interface{}
// 再分别添加变量，最后进行序列化
func (c *ClassRoom) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.students)
}

func (c *ClassRoom) UnmarshalJSON(buf []byte) error {
	return json.Unmarshal(buf, &c.students)
}

func (c *ClassRoom) List() error {
	for stu := range c.students {
		fmt.Println(stu, c.students[stu])
	}
	return nil
}
func (c *ClassRoom) Update(name string, id int) error {
	if _, ok := c.students[name]; !ok {
		return fmt.Errorf("not exist info for %s", name)
	}
	c.students[name].id = id
	return nil
}

func (c *ClassRoom) Add(name string, id int) error {
	if _, ok := c.students[name]; ok {
		return fmt.Errorf("duplicated info")
	}
	c.students[name] = &Student{
		name: name,
		id:   id,
	}
	return nil
}

func save(classRoom map[string]*ClassRoom, fileName string) error {
	buf, err := json.Marshal(classRoom)
	if err != nil {
		return nil
	}
	fd, err := os.Open(fileName)
	if err != nil {
		return nil
	}
	defer fd.Close()
	fd.Write(buf)
	return nil
}

func main() {
	crs := make(map[string]*ClassRoom)
	fmt.Println("crs:", crs)
	cr1 := &ClassRoom{
		students: make(map[string]*Student),
	}
	cr1.Add("binggan", 1)
	cr1.List()

	crs["51reboot"] = cr1
	cr1.Update("binggan", 3)
	cr1.List()
	fmt.Println("crs:", crs)
	save(crs, "a.txt")
}
