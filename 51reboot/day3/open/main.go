package main

import (
	"log"
	"os"
)

func read1() {
	f, err := os.Create("a.txt")
	if err != nil {
		log.Fatal(err)
	}
	f.WriteString("Hello\n")
	f.Close()
}

func read2() {
	f, err := os.OpenFile("a.txt", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0755)
	if err != nil {
		log.Fatal(err)
	}
	f.WriteString("world\n")
	f.Close()
}

func read3() {
	f, err := os.OpenFile("a.txt", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0755)
	if err != nil {
		log.Fatal(err)
	}

	f.Seek(0, 0)
	f.WriteString("$$$$$$$\n")
	f.Close()
}

func main() {
	read1()
	read2()
	read3()
}
