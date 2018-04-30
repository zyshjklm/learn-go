package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/jungle85gopy/learn-go/queue2worker/common"
)

const addr = "http://127.0.0.1:8080/upload"

func genData(metric string) common.DataInfo {
	// modify you process logic here

	di := common.DataInfo{Metric: metric}

	di.Timestamp = time.Now().Unix()
	di.Value = int64(rand.Intn(1000))
	return di
}

func main() {
	tickerSend := time.NewTicker(time.Millisecond * 100)
	for {
		select {
		case <-tickerSend.C:
			data := &common.DataCollection{
				Version: "1.0",
				Token:   "TOKEN12345678",
			}
			data.DataSlice = append(data.DataSlice, genData("cpu"))
			data.DataSlice = append(data.DataSlice, genData("mem"))

			buf, _ := json.Marshal(data)
			str := string(buf)
			log.Print(str)
			resp, err := http.Post(addr, "application/json", strings.NewReader(str))
			if err != nil {
				log.Println("err resp:", err)
				time.Sleep(time.Second * 1)
				continue
			}
			defer resp.Body.Close()

			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Println("err body:", err)
			}
			if len(body) > 0 {
				fmt.Println(string(body))
			}
		}
	}
}
