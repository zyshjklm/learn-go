package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

var cmdout *exec.Cmd

func main() {
	var cmd, params string
	f := bufio.NewReader(os.Stdin)
	for {
		cmd, params = "", ""
		fmt.Print("\ninput cmd>")
		line, _ := f.ReadString('\n')
		if len(line) == 1 {
			continue
		}
		line = strings.Replace(line, "\n", "", -1)
		// fmt.Println("your input:", line)
		fmt.Sscan(line, &cmd, &params)
		if cmd == "stop" {
			break
		}
		// fmt.Printf("cmd:%s,params:%s\n", cmd, params)
		if len(params) == 0 {
			cmdout = exec.Command(cmd)
		} else {
			cmdout = exec.Command(cmd, params)
		}
		out, _ := cmdout.StdoutPipe()
		if err := cmdout.Start(); err != nil {
			log.Fatal(err)
		}
		frd := bufio.NewReader(out)
		// fmt.Println("-- start read --")
		for {
			iline, err := frd.ReadString('\n')
			if err != nil {
				// fmt.Printf("read err:%v", err)
				break
			}
			fmt.Print(iline)
		}
		cmdout.Wait()
		// 如果不定义这一行就会产生一个僵尸进程，
		// 可以通过该参数来获取当前程序的终止状态，
		// 可以告知程序员当前程序是如何终止的。
	}
}
