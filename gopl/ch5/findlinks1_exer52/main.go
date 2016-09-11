package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}
	
	count := map[string]uint32{}
	elementCount(count, doc)
	for k, v := range count {
		fmt.Println(k, v)
	}
}

func elementCount(cnt map[string]uint32, node *html.Node) {
	if node.Type == html.ElementNode {
		cnt[node.Data]++ 	// tag ++
		// fmt.Println(cnt)
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		elementCount(cnt, c)
	}
}

