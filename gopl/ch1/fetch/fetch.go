// Fetch and print the content found at a URL.
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"io"
	"strings"
)


func fetch() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		body, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", body)
	}
}

// Exercise 1.7
// Exercise 1.8
// Exercise 1.9

func fetchExercise() {
	prefix := "http://"

	for _, url := range os.Args[1:] {

		// exercise 1.8
		if !strings.HasPrefix(url, prefix) {
			url = prefix + url
		}

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		num, err := io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: copy %s to Stdout: %v\n", url, err)
			os.Exit(1)
		} else {
			fmt.Fprintf(os.Stdout, "\n\nfetch: status\n%s", resp.Status)
			// 200 OK
			fmt.Printf("\n\ntotal length: %d\n", num)
		}
	}
}

func main() {
	//fetch()
	fetchExercise()
}
