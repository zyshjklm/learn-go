package main

import (
	"encoding/json"
	"io"
	"log"
	"strings"
)

type User struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

var jsonString string = `{
    "username":"rsj217@gmail.com",
    "password":"123"
}`

func MyDecode(r io.Reader) (u *User, err error) {
	u = new(User)
	err = json.NewDecoder(r).Decode(u)
	return
}

func main() {
	user, err := MyDecode(strings.NewReader(jsonString))
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("%#v\n", user)
}
