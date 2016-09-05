package main  

import (
	"fmt"
)

var graph = make(map[string]map[string]bool)

func addEdge(from, to string) {
	edges := graph[from]

	if edges == nil {
		edges = make(map[string]bool)
		graph[from] = edges
	}
	edges[to] = true
}

func hasEdge(from, to string) bool {
	return graph[from][to]
}

func main() {
	f, t := "dotA", "dotB"

	fmt.Println(graph)
	fmt.Println(hasEdge(f, t))

	addEdge(f, t)

	fmt.Println(graph)
	fmt.Println(hasEdge(f, t))


	addEdge(f, "CCC")

	fmt.Println(graph)
	fmt.Println(hasEdge(f, t))

}