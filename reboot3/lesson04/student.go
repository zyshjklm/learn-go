package main

import (
	"bufio"
	"fmt"
	"os"
)

// Student for student
type Student struct {
	ID   int
	Name string
}

// name should be uniq
var students = make(map[string]Student)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("> ")
		scanner.Scan()
		line := scanner.Text()
		var cmd string
		fmt.Scan(line, &cmd)
		fmt.Printf("cmd:%s\n", cmd)
		switch cmd {
		case "add":
			// add id name
			var id int
			var name string
			fmt.Scan(line, &id, &name)
			students[name] = Student{ID: id, Name: name}
		case "list":
			for k, v := range students {
				fmt.Printf("%d %s\n", k, v)
			}
		}
	}
}
