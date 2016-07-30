// Fetch all URLs in parallel and 
// reports their times and sizes
package main 

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
	"strings"
)

const prefix = "http://"

func main() {
	start := time.Now()
	ch := make(chan string)

	for _, url := range os.Args[1:] {
		go fetchall(url, ch)	// start a goroutine
	}

	for range os.Args[1:] {
		fmt.Println(<- ch)	// receive from channel ch
	}

	fmt.Printf("%.4fs elapsed\n", time.Since(start).Seconds())
}

func fetchall(url string, ch chan<- string) {
	start := time.Now()
	if !strings.HasPrefix(url, prefix) {
			url = prefix + url
	}
	resp, err := http.Get(url)

	if err != nil {
		// process err
		ch <- fmt.Sprint(err)
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()	// don't leak resources

	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v\n", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.4fs %7d %s", secs, nbytes, url)
}


