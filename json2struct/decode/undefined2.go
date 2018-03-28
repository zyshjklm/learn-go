package main

import (
	"encoding/json"
	"fmt"
)

var jsonString string = `{
	"things": [
		{
			"name": "Alice",
			"age": 37
		},
		{
			"city": "Ipoh",
			"country": "Malaysia"
		},
		{
			"name": "Bob",
			"age": 36
		},
		{
			"city": "Northampton",
			"country": "England"
		}
	]
}`

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Place struct {
	City    string `json:"city"`
	Country string `json:"country"`
}

type Mixed struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	City    string `json:"city"`
	Country string `json:"country"`
}

func decode(jsonStr []byte) (persons []Person, places []Place) {
	var data map[string][]Mixed
	err := json.Unmarshal(jsonStr, &data)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%#v\n", data["things"])

	for i := range data["things"] {
		item := data["things"][i]
		if item.Name != "" {
			persons = append(persons, Person{Name: item.Name, Age: item.Age})
		} else {
			places = append(places, Place{City: item.City, Country: item.Country})
		}
	}
	return
}

func main() {
	personsA, placesA := decode([]byte(jsonString))
	fmt.Printf("%#v\n", personsA)
	fmt.Printf("%#v\n", placesA)
}
