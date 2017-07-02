package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	fmt.Println(os.Stdin)
	fmt.Println(os.Stdout)
	fmt.Println(os.Stderr)

	os.Stdout.WriteString("hello golang\n")
	os.Stderr.WriteString("hello stderr\n")

	f, err := os.Create("ls.out")
	if err != nil {
		log.Fatal(err)
	}
	cmd := exec.Command("ls", "-l")
	cmd.Stdout = f
	cmd.Start()
	cmd.Wait()
}
