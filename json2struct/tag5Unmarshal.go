package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// Account struct
type Account struct {
	Email    string  `json:"email"`
	Password string  `json:"password"`
	Money    float64 `json:"money,string"`
}

func main() {
	account := Account{}
	var jsonString string = `{
		"email":"rsj217@gmail.com",
		"password":"123456",
		"money":"100.5"
	}`

	err := json.Unmarshal([]byte(jsonString), &account)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%+v\n", account)
}

/*
go run tag4String.go
{"email":"rsj217@gmail.com","money":"100.5"}
*/
