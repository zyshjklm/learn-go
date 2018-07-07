package main

import (
	"encoding/json"
	"fmt"
)

type S struct {
	DiskIOPS string `json:"diskiops"`
}

type D struct {
	DiskIOPS string `db:"disk-IOPS"`
}

func main() {
	var d = D{
		DiskIOPS: "100",
	}
	var s = S{}

	jd, _ := json.Marshal(&d)
	fmt.Println(string(jd))

	if err := json.Unmarshal(jd, &s); err != nil {
		panic(err)
	}
	fmt.Printf("d:%+v\n", d)
	fmt.Printf("s:%+v\n", s)
}
