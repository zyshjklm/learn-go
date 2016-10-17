package main 

import (
	"fmt"
	"encoding/json"
)

type Message struct {
	Name string
	Body string
	Time int64
	secret string
	// secret will not decoding in json.
}

func encode_from_struct(msg Message) []byte {
	// encoding
	js, err := json.Marshal(msg)
	if err != nil {
		fmt.Println("Marshal error!")
	}
	fmt.Println(js)
	return js
}

func decode_to_struct(msg []byte) Message {
	// decoding
	var newMsg Message
	err := json.Unmarshal(msg, &newMsg)
	if err != nil {
		fmt.Println("Unmarshal error!")
	}

	return newMsg
}

func decode_byte() {
	
}

func main() {
	var js []byte
	var msg1, msg2 Message
	
	msg1 = Message{"Alice", "Hello", 1294706395881547000, "bj"}
	fmt.Println("origin:", msg1)

	js = encode_from_struct(msg1)
	msg2 = decode_to_struct(js)
	
	// no "local" key here.
	fmt.Println("new:", msg2)

	// decoding from other byte
	var b []byte = []byte(`{"Name":"Bob","Food":"Pickle"}`)
	msg2 = decode_to_struct(b)
	fmt.Println("from byte:", msg2)

}

