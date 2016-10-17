package main 

import (
	"encoding/json"
	"log"
	"os"
)

type Message struct {
	Name string
	Body string
	Time int64
	secret string
	// secret will not decoding in json.
}


func main() {
	dec := json.NewDecoder(os.Stdin)
	enc := json.NewEncoder(os.Stdout)

	for {
		var v map[string]interface{}
		if err := dec.Decode(&v); err != nil {
			log.Fatalln(err)
			// return
		}
		for k := range v {
			if k != "Name" {
				delete(v, k)	// delete except Name field.
			}
		}

		if err := enc.Encode(&v); err != nil {
			log.Println(err)
		}
	}

}

