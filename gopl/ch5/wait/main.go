package main

import (
	"fmt"
	"time"
	"log"
	"os"
	"net/http"
)


// attempts to contact the server of a URL.
// It tires for one minute using exponential back-off.
// It reports an error if all attempts fail.
func WaitForServer(url string) error {
	const timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout)

	log.SetPrefix("[wait] ")
	//log.SetFlags(0)
	for tries := 0; time.Now().Before(deadline); tries++ {
		log.Printf("Head url %s", url)
		_, err := http.Head(url)
		if err == nil {
			return nil 		// success
		}
		log.Printf("server not responding (%s); retrying...", err)
		// exponential back-off.
		time.Sleep(time.Second << uint(tries))
	}
	log.SetPrefix("")
	return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: wait url\n")
		os.Exit(1)
	}
	url := os.Args[1]
	if err := WaitForServer(url); err != nil {
		log.Fatalf("Site is down: %v\n", err)
		os.Exit(1)
	}
}

