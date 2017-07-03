package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	host, _ := os.Hostname()
	prompt := fmt.Sprintf("[jungle@%s]$ ", host)
	r := bufio.NewScanner(os.Stdin)
	// r := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(prompt)
		if !r.Scan() {
			fmt.Println(r.Err())
			break
		}
		fmt.Println("\t--r.Err():", r.Err())
		line := r.Text()

		// 注释的部分，使用了另一种方式来获取输入
		// line, _ := r.ReadString('\n')
		// line = strings.TrimSpace(line)

		if len(line) == 0 {
			continue
		}
		args := strings.Fields(line)
		cmd := exec.Command(args[0], args[1:]...)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			fmt.Println(err)
		}
	}
}
