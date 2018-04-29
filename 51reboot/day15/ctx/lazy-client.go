package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

var (
	wg sync.WaitGroup
	to time.Duration = 2 // timeout
)

// ResPack for resp
type ResPack struct {
	r   *http.Response
	err error
}

func work(ctx context.Context) {
	tr := &http.Transport{}
	client := &http.Client{Transport: tr}
	defer wg.Done()
	c := make(chan ResPack, 1)

	req, _ := http.NewRequest("GET", "http://localhost:8021", nil)
	go func() {
		resp, err := client.Do(req)
		pack := ResPack{r: resp, err: err}
		c <- pack
	}()

	select {
	case <-ctx.Done():
		tr.CancelRequest(req)
		<-c
		fmt.Println("Timeout!")
	case res := <-c:
		if res.err != nil {
			fmt.Println(res.err)
			return
		}
		defer res.r.Body.Close()
		out, _ := ioutil.ReadAll(res.r.Body)
		fmt.Printf("Server Response: %s", out)
	}
	return
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), to*time.Second)
	defer cancel()
	wg.Add(1)
	go work(ctx)
	wg.Wait()
	fmt.Println("Finished")
}
