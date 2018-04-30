package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/jungle85gopy/learn-go/queue2worker/common"
)

// global variable
var (
	MaxLen    int // MaxLen max len of post data
	MaxWorker int // MaxWorker max worker to consume queue
	MaxQueue  int // MaxQueue max chan queue for client post

	dispatcher *common.Dispatcher
)

func init() {
	MaxWorker, _ = strconv.Atoi(os.Getenv("MAX_WORKER"))
	MaxQueue, _ = strconv.Atoi(os.Getenv("MAX_QUEUE"))
	MaxLen = 1024 * 1024
}

// DataHandler handler for data
func DataHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	// Read the body into a string for json decoding
	var content = &common.DataCollection{}
	err := json.NewDecoder(io.LimitReader(r.Body, int64(MaxLen))).Decode(&content)
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// Go through each data and queue items individually to be posted
	for _, data := range content.DataSlice {
		// create a job with the DataInfo, and push the job into the queue.
		dispatcher.JobQueue <- common.Job{DataInfo: data}
		// dispatcher.JobQueue <- job
	}
	w.WriteHeader(http.StatusOK)
}

func main() {
	log.Printf("queue:%d, worker:%d\n", MaxQueue, MaxWorker)

	dispatcher = common.NewDispatcher(MaxQueue, MaxWorker)
	dispatcher.Run()

	http.HandleFunc("/upload", DataHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

/*
	Usage:
	MAX_QUEUE=1024 MAX_WORKER=20 go run server.go
*/
