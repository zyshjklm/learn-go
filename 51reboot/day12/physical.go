package main

import (
	"log"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
)

func main() {
	for {
		// cpu
		core, err := cpu.Percent(time.Second, false)
		if err != nil {
			panic(err)
		}
		log.Println("single cpu:", core)

		cores, err := cpu.Percent(time.Second, true)
		if err != nil {
			panic(err)
		}
		log.Println("cpu cores:", cores)

		// mem
		memstat, err := mem.VirtualMemory()
		if err != nil {
			panic(err)
		}
		log.Println("mem percent:", memstat.UsedPercent)
		log.Printf("mem used: %dM\n", memstat.Used/1024/1024)

		// disk
		for _, d := range []string{"/", "/home"} {
			diskstat, err := disk.Usage(d)
			if err != nil {
				panic(err)
			}
			log.Printf("%s used percent:%.2f\n", d, diskstat.UsedPercent)
		}

		// load
		loadavg, err := load.Avg()
		if err != nil {
			panic(err)
		}
		log.Println(loadavg)

		log.Println("----- end loops ----\n")
	}
}
