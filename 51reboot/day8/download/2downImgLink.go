package main

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func repairLink(uri *url.URL, url string) string {
	switch {
	case strings.HasPrefix(url, "http"):
		return url
	case strings.HasPrefix(url, "//"):
		return uri.Scheme + ":" + url
	case strings.HasPrefix(url, "/"):
		return uri.Scheme + "://" + uri.Host + url
	default:
		// golang-spider/xyz/img.html
		base := filepath.Dir(uri.Path)
		return uri.Scheme + "://" + uri.Host + base + "/" + url
	}
}

func cleanLinks(oriURL string, links []string) ([]string, error) {
	var result []string
	uri, err := url.Parse(oriURL)
	if err != nil {
		return nil, err
	}
	for _, url := range links {
		result = append(result, repairLink(uri, url))
	}
	return result, nil
}

func fetchOrigin(url string) ([]string, error) {
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
	return cleanLinks(url, urls)
}

func downloadImgs(urls []string, tmpDir string) error {
	for _, url := range urls {
		resp, err := http.Get(url)
		if err != nil {
			return err
		}
		// close resp
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			return errors.New(resp.Status)
		}
		name := path.Base(url)
		fullName := filepath.Join(tmpDir, name)
		fmt.Println("dest:", fullName)
		fd, err := os.Create(fullName)
		if err != nil {
			return err
		}
		defer fd.Close()
		_, err = io.Copy(fd, resp.Body)
		if err != nil {
			log.Println(err)
		}
	}
	return nil
}

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("usage: %s url", os.Args[0])
		os.Exit(1)
	}
	url := os.Args[1]
	// "http://59.110.12.72:7070/golang-spider/img.html"
	links, err := fetchOrigin(url)
	if err != nil {
		fmt.Println(err)
	}
	for i, url := range links {
		fmt.Println(i, url)
	}
	tmpDir, err := ioutil.TempDir("", "spider")
	if err != nil {
		log.Fatal(err)
	}
	// defer os.RemoveAll(tmpDir)
	err = downloadImgs(links, tmpDir)
	if err != nil {
		log.Panic(err)
	}
}

/*
dest: /var/folders/4_/20sh50zs0lg64gphvscqq2gm0000gn/T/spider218105315/v2-58e318de6172810c1b3c7236e8e0dbb4.jpg
dest: /var/folders/4_/20sh50zs0lg64gphvscqq2gm0000gn/T/spider218105315/v2-40becd4a519329198ecb3807f342fd7b.jpg
dest: /var/folders/4_/20sh50zs0lg64gphvscqq2gm0000gn/T/spider218105315/a.jpg
dest: /var/folders/4_/20sh50zs0lg64gphvscqq2gm0000gn/T/spider218105315/b.jpg
*/
