package common

import (
	"log"
)

// ------- modify your data here : start -------- //

// DataInfo info
type DataInfo struct {
	Metric    string `json:"metric"`
	Value     int64  `json:"value"`
	Timestamp int64  `json:"time"`
}

// ProcessData process data
func (p *DataInfo) ProcessData() error {
	// process logic here

	log.Println("[process]:", p)

	return nil
}

// ------- modify your data here : end -------- //

// DataCollection for collect DataInfo from client
type DataCollection struct {
	Version   string     `json:"version"`
	Token     string     `json:"token"`
	DataSlice []DataInfo `json:"data"`
}
