package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// Account struct
type Account struct {
	Email    string
	password string
	Money    float64
}

// User struct
type User struct {
	Name    string
	Age     int
	Roles   []string
	Skill   map[string]float64
	Account Account
	Extra   []interface{}
	Level   map[string]interface{}
}

func main() {
	// 字符串有引号，浮点数没有引号
	account := Account{
		Email:    "rsj217@gmail.com",
		password: "123456",
		Money:    100.5,
	}

	skill := make(map[string]float64)
	skill["python"] = 99.5
	skill["elixir"] = 90
	skill["ruby"] = 80.0

	level := make(map[string]interface{})

	level["web"] = "Good"
	level["server"] = 90
	level["tool"] = nil

	user := User{
		Name:    "rsj217",
		Age:     27,
		Roles:   []string{"Owner", "Master"},
		Skill:   skill,
		Account: account,
		Level:   level,
	}

	rs, err := json.Marshal(user)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(rs))
}

/*
go run base5mapInterface.go
{"Name":"rsj217","Age":27,"Roles":["Owner","Master"],"Skill":{"elixir":90,"python":99.5,"ruby":80},"Account":{"Email":"rsj217@gmail.com","Money":100.5},"Extra":null,"Level":{"server":90,"tool":null,"web":"Good"}}
*/
