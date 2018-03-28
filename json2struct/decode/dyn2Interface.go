package main

import (
	"encoding/json"
	"io"
	"log"
	"strings"
)

type User struct {
	UserName interface{} `json:"username"`
	Password string      `json:"password"`
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
	err = json.NewDecoder(r).Decode(u)

	// 使用断言来解决不同类型的取值
	switch t := u.UserName.(type) {
	case string:
		u.UserName = t
	case float64:
		u.UserName = int64(t)
	}
	return
}

func main() {
	user, _ := MyDecode(strings.NewReader(jsonString))
	log.Printf("%#v\n", user)

	user, _ = MyDecode(strings.NewReader(jsonString2))
	log.Printf("%#v\n", user)
}
