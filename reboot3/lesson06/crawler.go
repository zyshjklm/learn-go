package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	// https://www.itlipeng.cn/2017/04/25/goquery-%E6%96%87%E6%A1%A3/
)

// crawl the given link, follow *.go file and path
func main() {
	if len(os.Args) != 2 {
		log.Fatalf("usage: %s url", os.Args[0])
	}

	crawl(os.Args[1])
}

func getResponse(url string) io.Reader {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	//defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", resp.StatusCode, resp.Status)
	}
	return resp.Body
}

func getFile(name string) io.Reader {
	fd, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	return fd
}

func crawl(baseURL string) {
	fmt.Printf("\nbase:%s\n", baseURL)
	resp := getResponse(baseURL)
	//resp := getFile(baseURL)
	fmt.Printf("  -- start parse : %s --\n", baseURL)
	url, err := url.Parse(baseURL)
	if err != nil {
		log.Fatal(err)
	}
	pwd, _ := os.Getwd()
	uri := url.RequestURI()
	fmt.Printf("  uri:%s\n", uri)
	if strings.HasPrefix(uri, "/") {
		os.MkdirAll(pwd+url.RequestURI(), os.ModePerm)
	}
	httpPrefix := fmt.Sprintf("%s://%s:%s", url.Scheme, url.Hostname(), url.Port())
	//fmt.Printf("  %s\n", httpPrefix)

	doc, err := goquery.NewDocumentFromReader(resp)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("dd").Find("a").Each(func(i int, a *goquery.Selection) {
		//text := a.Text()
		val, _ := a.Attr("href")
		//fmt.Printf("\thref:%s\n", val)
		//fmt.Printf("\ttext:%s\n", text)

		if strings.HasSuffix(val, ".go") || strings.HasSuffix(val, "main.slide") {
			fmt.Printf("\tGet:%s\n", httpPrefix+val)
			gores, goerr := http.Get(httpPrefix + val)
			if goerr != nil {
				fmt.Printf("get %s err\n", val)
			}
			fd, err := os.Create(pwd + val)
			if err != nil {
				log.Print(err)
			}
			io.Copy(fd, gores.Body)
		} else {
			// process path
			crawl(httpPrefix + val)
		}
	})
}
