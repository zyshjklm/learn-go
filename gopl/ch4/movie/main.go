package main 

import (
	"fmt"
	"log"
	"encoding/json"
)

type movie struct {
	Title string
	Year int `json:"released"`
	Color bool `json:"color,omitempty"`
	Actors []string
}

func main() {
	var movies = []movie{
		{Title: "Casablanca", Year: 1942, Color: false,
			Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
		{Title: "Cool Hand Luke", Year: 1967, Color: true,
			Actors: []string{"Paul Newman"}},
		{Title: "Bullitt", Year: 1968, Color: true,
			Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
	}

	data1, err := json.Marshal(movies)
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s\n", err)
	}

	fmt.Printf("%s\n", data1)

	data2, err := json.MarshalIndent(movies, "", "    ")
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s\n", err)
	}

	fmt.Printf("%s\n", data2)
}