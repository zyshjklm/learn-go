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

	cur, _ := f.Seek(0, 0)
	log.Println("current index:", cur)
	// 0; but write to the end of file, because open a.txt with O_APPEND
	f.WriteString("$$$$$$$\n")
	f.Close()
}

func read4() {
	f, err := os.OpenFile("a.txt", os.O_CREATE|os.O_RDWR, 0755)
	if err != nil {
		log.Fatal(err)
	}
	cur, _ := f.Seek(0, 2)
	log.Println("current index:", cur)
	f.WriteString("EEEEEEEE\n")

	cur, _ = f.Seek(0, 0)
	log.Println("current index:", cur)
	f.WriteString("***")
	f.Close()
	/*
	   ***lo
	   world
	   $$$$$$$
	   EEEEEEEE
	*/
}

func main() {
	read1()
	read2()
	read3()
	read4()
}
