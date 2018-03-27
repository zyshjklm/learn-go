package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// Account struct
type Account struct {
	Email    string  `json:"email"`
	Password string  `json:"pass_word"`
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
go run tag1Base.go
{"email":"rsj217@gmail.com","pass_word":"123456","money":100.5}
*/
