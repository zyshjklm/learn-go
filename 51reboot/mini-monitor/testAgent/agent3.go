package main

import (
	"encoding/json"
	"flag"
	"log"
	"net"
	"os"
	"runtime"
	"time"

	"github.com/jungle85gopy/learn-go/51reboot/mini-monitor/common"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
)

var transAddr = flag.String("trans", ":6000", "transfer address")

// NewMetric new a Metric
func NewMetric(metric string, value float64) *common.Metric {
	hostname, _ := os.Hostname()
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

	var metric *common.Metric
	var err error
	conn, err := net.Dial("tcp", *transAddr)
	if err != nil {
		panic(err)
	}

	ticker := time.NewTicker(time.Second * 1)

	for range ticker.C {
		cpus, err := cpu.Percent(time.Second, false)
		if err != nil {
			panic(err)
		}
		metric = NewMetric("cpu.usage", cpus[0])
		buf, _ := json.Marshal(metric)
		_, err = conn.Write(buf)
		if err != nil {
			log.Printf("writer err: %s", err.Error())
		}
		log.Println(string(buf))

		mem, err := mem.VirtualMemory()
		if err != nil {
			log.Print(err)
		}
		metric = NewMetric("mem.usage", float64(mem.Total)*mem.UsedPercent)
		buf, _ = json.Marshal(metric)
		_, err = conn.Write(buf)
		if err != nil {
			log.Printf("writer err: %s", err.Error())
		}
		log.Println(string(buf))
	}
}
