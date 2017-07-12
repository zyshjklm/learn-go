package main

import (
	"errors"
	"fmt"
)

func main() {
	var e error
	e = errors.New("an error")
	fmt.Println(e.Error())
	fmt.Println(e)

	e = fmt.Errorf("err from fmt.Errorf")
	fmt.Println(e)
}

/*
http://localhost:6060/pkg/builtin/

type error interface {
    Error() string
}
*/
