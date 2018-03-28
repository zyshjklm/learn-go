package main

import (
	"encoding/json"
	"io"
	"log"
	"strings"
)

type User struct {
	UserName json.RawMessage `json:"username"`
	Password string          `json:"password"`

	Email string
	Phone int64
}

var jsonString string = `{
    "username":"rsj217@gmail.com",
    "password":"123"
}`

var jsonString2 string = `{
    "username":18612349876,
    "password":"123"
}`

func MyDecode(r io.Reader) (u *User, err error) {
	u = new(User)
	if err = json.NewDecoder(r).Decode(u); err != nil {
		return
	}
	var email string
	if err = json.Unmarshal(u.UserName, &email); err == nil {
		u.Email = email
		return
	}
	var phone int64
	if err = json.Unmarshal(u.UserName, &phone); err == nil {
		u.Phone = phone
	}
	return
}

func main() {
	user, _ := MyDecode(strings.NewReader(jsonString))
	log.Printf("%#v\n", user)

	user, _ = MyDecode(strings.NewReader(jsonString2))
	log.Printf("%#v\n", user)
}
