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

func main() {
	// 字符串有引号，浮点数没有引号
	account := Account{
		Email:    "rsj217@gmail.com",
		password: "123456",
		Money:    100.5,
	}

	rs, err := json.Marshal(account)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(rs)
	fmt.Println(string(rs))
}

/*
go run 1marshal.go
[123 34 69 109 97 105 108 34 58 34 114 115 106 50 49 55 64 103 109 97 105 108 46 99 111 109 34 44 34 77 111 110 101 121 34 58 49 48 48 46 53 125]
{"Email":"rsj217@gmail.com","Money":100.5}
*/
