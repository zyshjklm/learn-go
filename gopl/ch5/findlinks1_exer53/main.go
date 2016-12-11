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
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}

	printCont(doc)
}

func printCont(node *html.Node) {
	if node.Type == html.ElementNode && node.Data == "script" {
		return
	}
	if node.Type == html.TextNode && node.Data != "script" {
		data := strings.TrimSpace(node.Data)
		if len(data) > 0 {
			fmt.Println(data)
		}
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		printCont(c)
	}
}

