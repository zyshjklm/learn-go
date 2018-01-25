package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	for _, name := range os.Args[1:] {
		f, err := os.Open(name)
		if err != nil {
			log.Fatal(err)
		}
		infos, err := f.Readdir(-1)
		if err != nil {
			fmt.Println(name, "not dir")
			continue
		}
		fmt.Printf("--- name: %s ---\n", name)
		for _, info := range infos {
			fmt.Printf("%v %d %s\n", info.IsDir(), info.Size(), info.Name())
		}
	}
}
