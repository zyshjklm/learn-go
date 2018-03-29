package main

import (
	"encoding/json"
	"fmt"
)

type Message struct {
	Asks [][]int `json:"Asks"`
	Bids [][]int `json:"Bids"`
}

func main() {
	b := []byte(`{"Asks": [[21, 1], [22, 1]] ,"Bids": [[20, 1], [19, 1]]}`)
	var m Message
	if err := json.Unmarshal(b, &m); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", m)
}
