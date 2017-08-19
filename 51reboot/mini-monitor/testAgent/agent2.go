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

func main() {
	conn, err := net.Dial("tcp", *transAddr)
	if err != nil {
		panic(err)
	}

	hostname, _ := os.Hostname()
	var metric []*common.Metric

	for {
		metric = nil
		cpus, err := cpu.Percent(time.Second, false)
		if err != nil {
			panic(err)
		}
		metricCPU := &common.Metric{
			Metric:    "cpu.usage",
			Endpoint:  hostname,
			Value:     cpus[0],
			Tag:       []string{runtime.GOOS},
			Timestamp: time.Now().Unix(),
		}
		metric = append(metric, metricCPU)

		mem, err := mem.VirtualMemory()
		if err != nil {
			log.Print(err)
		}
		metricMem := &common.Metric{
			Metric:    "mem.usage",
			Endpoint:  hostname,
			Value:     float64(mem.Total) * mem.UsedPercent,
			Tag:       []string{runtime.GOOS},
			Timestamp: time.Now().Unix(),
		}
		metric = append(metric, metricMem)

		buf, _ := json.Marshal(metric)
		n, err := conn.Write(buf)
		if err != nil {
			log.Printf("writer err: %s", err.Error())
		}
		log.Printf("write len:%d, buf len:%d,  status:%v\n", n, len(buf), n == len(buf))
		log.Println(string(buf) + "\n")
	}
}
