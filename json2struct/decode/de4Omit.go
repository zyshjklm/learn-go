package main

import (
	"encoding/json"
	"fmt"
)

type Account struct {
	Email    string  `json:"email"`
	Password string  `json:"password"`
	Money    float64 `json:"-"`
}

var jsonString string = `{
    "email":"rsj217@gmail.com",
    "password":"123",
    "money":100.5
}`

func main() {
	account := Account{}
	_ = json.Unmarshal([]byte(jsonString), &account)

	fmt.Printf("%+v\n", account)
}
