package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// User struct
type User struct {
	Name  string
	Age   int
	Roles []string
	Extra []interface{}
}

func main() {
	extra := []interface{}{123, "hello golang"}
	user := User{
		Name:  "rsj217",
		Age:   27,
		Roles: []string{"Owner", "Master"},
		Extra: extra,
	}

	rs, err := json.Marshal(user)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(rs))
}

/*
go run 4diffArray.go
{"Name":"rsj217","Age":27,"Roles":["Owner","Master"],"Extra":[123,"hello golang"]}
*/
