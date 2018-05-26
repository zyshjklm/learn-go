package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type originData struct {
	Name string
	Tag  string
	Type interface{}
}

var docStruct map[int]originData

/*
OrgID   org_id  ""
Name    name    ""
City    city    ""
ID      id      0
Yidu    yidu    false
*/
func readDocs() {
	fd, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	// scanner := bufio.NewScanner(fd)
	reader := bufio.NewReader(fd)
	for {
		// scanner.Scan()
		// line := scanner.Text()
		line, _ := reader.ReadString('\n')
		line = strings.TrimSuffix(line, "\n")
		if len(line) == 0 {
			break
		}
		var name, tag, nType string

		// log.Print(line)
		fmt.Sscan(line, &name, &tag, &nType)
		// fmt.Printf("type of name:%+v\n", reflect.TypeOf(nType))
		fmt.Println("name:", name, "tag:", tag, "type:", nType)
	}
	// buf, err := ioutil.ReadAll(fd)
	// fmt.Println(buf)
}
func main() {
	readDocs()
}
