package main

import (
	"encoding/json"
	"fmt"
)

type Account struct {
	Email    string  `json:"email"`
	PassWord string  `json:"password"`
	Money    float64 `json:"money"`
}

type Account2 struct {
	Email    string  `json:"email"`
	passWord string  `json:"password"`
	Money    float64 `json:"money"`
}
type Account3 struct {
	Email    string `json:"email"`
	Password string `json:"password"`
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

	account2 := Account2{}
	_ = json.Unmarshal([]byte(jsonString), &account2)
	fmt.Printf("%+v\n", account2)

	account3 := Account3{}
	_ = json.Unmarshal([]byte(jsonString), &account3)
	fmt.Printf("%+v\n", account3)
}
