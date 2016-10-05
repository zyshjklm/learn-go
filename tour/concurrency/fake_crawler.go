package main

import (
    "fmt"
    "sync"
)

type Fetcher interface {
    // Fetch returns the body of URL and
    // a slice of URLs found on that page.
    Fetch(url string) (body string, urls []string, err error)
}

type fakeResult struct {
    body string
    urls []string
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

func (f *fakeFetcher) Fetch(url string) (string, []string, error) {
    if res, ok := (*f)[url]; ok {
        return res.body, res.urls, nil
    }
    return "", nil, fmt.Errorf("not found: %s", url)
}


// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, 
           out chan string, end chan bool) {

    if depth <= 0 {
        end <- true
        return
    }

    // check current url been crawled
    if _, ok := crawled[url]; ok {
        out <- fmt.Sprintf("  -- ignore url: %s", url)
        end <- true
        return
    }
    // lock and modify status of current url
    crawledMutex.Lock()
    crawled[url] = true
    crawledMutex.Unlock()

    body, urls, err := fetcher.Fetch(url)
    if err != nil {
        out <- fmt.Sprintln(err)
        end <- true
        return
    }

    // normal return
    out <- fmt.Sprintf("found url: %s. \tbody: %q\n", url, body)

    // recursively crawl
    subEnd := make(chan bool)
    for _, u := range urls {
        go Crawl(u, depth-1, fetcher, out, subEnd)
    }

    for i := 0; i < len(urls); i++ {
        <- subEnd
    }
    end <- true
}



var crawled = make(map[string]bool)
var crawledMutex sync.Mutex


func main() {
    out, end := make(chan string), make(chan bool)

    // use pointer of fetcher to omit copy struct 
    go Crawl("http://golang.org/", 4, &fetcher, out, end)
    for {
        select {
        case result := <- out:
            fmt.Println(result)
        case <- end:
            fmt.Println("case end true...")
            return      // not break here.
        }
    }
}


// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
    "http://golang.org/": &fakeResult{
        "The Go Programming Language",
        []string{
            "http://golang.org/pkg/",
            "http://golang.org/cmd/",
        },
    },
    "http://golang.org/pkg/": &fakeResult{
        "Packages",
        []string{
            "http://golang.org/",
            "http://golang.org/cmd/",
            "http://golang.org/pkg/fmt/",
            "http://golang.org/pkg/os/",
        },
    },
    "http://golang.org/pkg/fmt/": &fakeResult{
        "Package fmt",
        []string{
            "http://golang.org/",
            "http://golang.org/pkg/",
        },
    },
    "http://golang.org/pkg/os/": &fakeResult{
        "Package os",
        []string{
            "http://golang.org/",
            "http://golang.org/pkg/",
        },
    },
}

