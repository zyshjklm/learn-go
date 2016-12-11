package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks: %v\n", err)
		os.Exit(1)
	}

	visit(doc)
}

func visit(n *html.Node) {
	fmt.Printf("Type: %d, Data: %s\n", n.Type, strings.TrimSpace(n.Data))

	for next := n.FirstChild; next != nil; next = next.NextSibling {
		visit(next)
	}
}

