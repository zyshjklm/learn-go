package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// Account struct
type Account struct {
	Email    string  `json:"email"`
	Password string  `json:"-"`
	Money    float64 `json:"money"`
}

func main() {
	// 字符串有引号，浮点数没有引号
	account := Account{
		Email:    "rsj217@gmail.com",
		Password: "123456",
		Money:    100.5,
	}

	rs, err := json.Marshal(account)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(rs))
}

/*
go run tag2Ignore.go
{"email":"rsj217@gmail.com","money":100.5}
*/
