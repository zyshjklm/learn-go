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
	"sync"

	"github.com/PuerkitoBio/goquery"
)

func repairLink(uri *url.URL, url string) string {
	switch {
	case strings.HasPrefix(url, "http"):
		return url
	case strings.HasPrefix(url, "//"):
		return fmt.Sprintf("%s:%s", uri.Scheme, url)
	case strings.HasPrefix(url, "/"):
		return fmt.Sprintf("%s://%s%s", uri.Scheme, uri.Host, url)
	default:
		urlPath := filepath.Dir(uri.Path) + "/" + url
		return fmt.Sprintf("%s://%s%s", uri.Scheme, uri.Host, urlPath)
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

// capsulate the waitgroup, concurrent pool and routine.
func downloadImgs(urls []string, dir string, conNum int) error {
	if conNum <= 0 {
		conNum = 1
	}
	var wg sync.WaitGroup
	wg.Add(conNum)
	urlChan := make(chan string)
	for i := 1; i <= conNum; i++ {
		go downWorker(urlChan, dir, &wg)
	}
	for _, url := range urls {
		urlChan <- url
	}
	close(urlChan)
	wg.Wait()
	return nil
}

// read from channel and call downloadURL()
func downWorker(urlChan chan string, dir string, wg *sync.WaitGroup) {
	for url := range urlChan {
		err := downloadURL(url, dir)
		if err != nil {
			log.Println(err)
		}
	}
	wg.Done()
}

func downloadURL(url, dir string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	// close resp
	defer resp.Body.Close()
	fullName := filepath.Join(dir, path.Base(url))
	fd, err := os.Create(fullName)
	if err != nil {
		return err
	}
	defer fd.Close()
	_, err = io.Copy(fd, resp.Body)
	if err != nil {
		return err
	}
	log.Println("down:", fullName)
	return nil
}

func fetchImages(w io.Writer, url string) {
	urls, err := fetchOrigin(url)
	if err != nil {
		log.Panic(err)
	}
	urls, _ = cleanLinks(url, urls)
	for i, url := range urls {
		fmt.Println(i, url)
	}

	tmpDir, err := ioutil.TempDir("", "spider")
	if err != nil {
		log.Panic(err)
	}
	// defer os.RemoveAll(tmpDir)
	err = downloadImgs(urls, tmpDir, 5)
	if err != nil {
		log.Panic(err)
	}
}

func handleHTTP(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	// 请求服务时，使用u=url.website.dom
	// "http://59.110.12.72:7070/golang-spider/img.html"
	url := req.FormValue("u")
	if len(url) != 0 {
		fetchImages(w, url)
	}
}

func main() {
	http.HandleFunc("/", handleHTTP)
	log.Fatal(http.ListenAndServe(":7071", nil))
}
