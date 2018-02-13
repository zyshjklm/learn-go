package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

func main() {
	baseDir := "/proc/"
	f, err := os.Open(baseDir)
	if err != nil {
		log.Fatal(err)
	}
	infos, _ := f.Readdir(-1)
	for _, info := range infos {
		if !info.IsDir() {
			// fmt.Printf("[cont] %v is not dir\n", info.Name())
			continue
		}
		pid, err := strconv.Atoi(info.Name())
		if err != nil {
			// fmt.Printf("[cont] strconv to int err:%v\n", err)
			continue
		}
		// dir and no err
		cmdLine := baseDir + info.Name() + "/cmdline"
		// fmt.Println("cmdline:", cmdLine)
		readFileName(pid, cmdLine)
	}
}

func readFileName(pid int, cmdLine string) {
	buf, err := ioutil.ReadFile(cmdLine)
	if err != nil {
		log.Fatal(err)
		return
	}
	fileName := string(buf)
	if len(fileName) > 0 {
		fmt.Printf("process: %d\tnamed: %s\n", pid, fileName)
	} else {
		fmt.Printf("process: %d\tbelong the SHELL\n", pid)
	}
}
