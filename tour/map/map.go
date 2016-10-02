package main 

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func map_init() {
	// style 1: var define and make create it.
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}	
	fmt.Println(m["Bell Labs"])

	// style 2: literals init
	var m2 = map[string]Vertex{
		"Bell Labs": Vertex{
			40.68433, -74.39967,
		},
		"Google": Vertex{
			37.42202, -122.08408,
		},
	}
	fmt.Println(m2)

	// 
	var m3 = map[string]Vertex{
		"Bell Labs": { 40.68433, -74.39967 },
		"Google": { 37.42202, -122.08408 },
	}
	fmt.Println(m3)
}

func map_matating() {
	m := make(map[string]int)
	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])
	
	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}

func main() {
	map_init()
	map_matating()
}