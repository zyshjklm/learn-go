package main

import (
	"flag"
	"log"
	"os"
	"runtime"
	"time"

	"github.com/jungle85gopy/learn-go/51reboot/mini-monitor/common"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
)

var transAddr = flag.String("trans", ":6000", "transfer address")
var hostname string

// NewMetric new a Metric
func NewMetric(metric string, value float64) *common.Metric {
	return &common.Metric{
		Metric:    metric,
		Endpoint:  hostname,
		Value:     value,
		Tag:       []string{runtime.GOOS},
		Timestamp: time.Now().Unix(),
	}
}

func main() {
	flag.Parse()
	hostname, _ = os.Hostname()

	sender := NewSender(*transAddr)
	ch := sender.Channel()
	go sender.Start()

	ticker := time.NewTicker(time.Second * 1)

	for range ticker.C {
		cpus, err := cpu.Percent(time.Second, false)
		if err != nil {
			panic(err)
		}
		metric := NewMetric("cpu.usage", cpus[0])
		ch <- metric

		mem, err := mem.VirtualMemory()
		if err != nil {
			log.Print(err)
		}
		metric = NewMetric("mem.usage", float64(mem.Total)*mem.UsedPercent/1024/1024)
		ch <- metric
	}
}
