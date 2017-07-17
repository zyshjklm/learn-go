package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

func fetch(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	// close resp
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return errors.New(resp.Status)
	}
	// display the html
	io.Copy(os.Stdout, resp.Body)
	return nil
}

func main() {
	url := os.Args[1]
	// url := "http://daily.zhidu.com"
	// url := "http://m.xiaohuar.com/"
	// url := "http://www.douban.com/"
	err := fetch(url)
	if err != nil {
		fmt.Println(err)
	}
}
