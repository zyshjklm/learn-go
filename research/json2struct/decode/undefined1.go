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

func decode(jsonStr []byte) (persons []Person, places []Place) {
	var data map[string][]map[string]interface{}
	err := json.Unmarshal(jsonStr, &data)
	if err != nil {
		fmt.Println(err)
		return
	}

	for i := range data["things"] {
		fmt.Println("i=:", i)
		item := data["things"][i]
		if item["name"] != nil {
			persons = addPerson(persons, item)
		} else {
			places = addPlace(places, item)
		}

	}
	return
}

func addPerson(persons []Person, item map[string]interface{}) []Person {
	name := item["name"].(string)
	age := item["age"].(float64)
	person := Person{name, int(age)}
	persons = append(persons, person)
	return persons
}

func addPlace(places []Place, item map[string]interface{}) []Place {
	city := item["city"].(string)
	country := item["country"].(string)
	place := Place{City: city, Country: country}
	places = append(places, place)
	return places
}

func main() {
	personsA, placesA := decode([]byte(jsonString))
	fmt.Printf("%#v\n", personsA)
	fmt.Printf("%#v\n", placesA)
}
