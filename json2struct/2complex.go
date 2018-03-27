package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// User struct
type User struct {
	Name  string
	Age   int
	Roles []string
	Skill map[string]float64
}

func main() {
	skill := make(map[string]float64)
	skill["python"] = 99.5
	skill["elixir"] = 90
	skill["ruby"] = 80.0

	user := User{
		Name:  "rsj217",
		Age:   27,
		Roles: []string{"Owner", "Master"},
		Skill: skill,
	}

	rs, err := json.Marshal(user)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(rs))
}

/*
go run 2complex.go
{"Name":"rsj217","Age":27,"Roles":["Owner","Master"],"Skill":{"elixir":90,"python":99.5,"ruby":80}}
*/
