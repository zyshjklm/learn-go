package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
)

func fetch(url string) ([]string, error) {
	var urls []string
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	// close resp
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}
	// display the html
	// io.Copy(os.Stdout, resp.Body)
	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		return nil, err
	}
	doc.Find("img").Each(func(i int, s *goquery.Selection) {
		link, ok := s.Attr("src")
		if ok {
			urls = append(urls, link)
		}
	})
	return urls, nil
}

func main() {
	url := os.Args[1]
	// url := "http://daily.zhidu.com"
	// url := "http://www.douban.com/"
	urls, err := fetch(url)
	if err != nil {
		fmt.Println(err)
	}
	for i, url := range urls {
		fmt.Println(i, url)
	}
}
