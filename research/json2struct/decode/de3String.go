package main

import (
	"encoding/json"
	"fmt"
)

type Account struct {
	Email    string  `json:"email"`
	Password string  `json:"password"`
	Money    float64 `json:"money,string"`
}

var jsonString string = `{
    "email":"rsj217@gmail.com",
    "password":"123",
    "money":"100.5"
}`

var jsonString2 string = `{
    "email":"rsj217@gmail.com",
    "password":"123",
    "money":100.95
}`

func main() {
	account := Account{}
	_ = json.Unmarshal([]byte(jsonString), &account)
	fmt.Printf("%#v\n", account)

	account2 := Account{}
	_ = json.Unmarshal([]byte(jsonString2), &account2)
	fmt.Printf("%#v\n", account2)
}
