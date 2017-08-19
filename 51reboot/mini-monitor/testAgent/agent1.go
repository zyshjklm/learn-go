package main

import (
	"encoding/json"
	"log"
	"os"
	"runtime"
	"time"

	"github.com/jungle85gopy/learn-go/51reboot/mini-monitor/common"
	"github.com/shirou/gopsutil/cpu"
)

func main() {
	hostname, _ := os.Hostname()
	for {
		cpus, err := cpu.Percent(time.Second, false)
		if err != nil {
			panic(err)
		}
		metric := &common.Metric{
			Metric:    "cpu.usage",
			Endpoint:  hostname,
			Value:     cpus[0],
			Tag:       []string{runtime.GOOS},
			Timestamp: time.Now().Unix(),
		}

		buf, _ := json.Marshal(metric)
		log.Println(string(buf))
	}
}
