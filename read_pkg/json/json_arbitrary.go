package main

import "fmt"
import "encoding/json"

func byte2json(b []byte) map[string]interface{} {
	var f map[string]interface{}
	err := json.Unmarshal(b, &f)
	if err != nil {
		fmt.Println("Unmarshal error!")
	} 	

	fmt.Println("data struct:", f)
	fmt.Println("\nformat style:")
	
	for k, v := range f {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case int:
			fmt.Println(k, "is int", vv)
		case []interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range vv {
				fmt.Println("  ", i, u)
			}
		}
	}

	return f
}

func main() {
	var b1 []byte
	var b2 map[string]interface{}

	b1 = []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)

 	b2 = byte2json(b1)

 	fmt.Println(b1)
 	fmt.Println(b2)
}

