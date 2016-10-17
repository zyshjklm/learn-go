package main

import (
	"bytes"
	"encoding/json"
	"log"
	"os"
)

type Road struct {
	Name string
	Number int
}


func main() {
	roads := []Road{
		{"Diamond Fork", 29},
		{"Sheep Creek", 51},
	}

	bt, err := json.Marshal(roads)
	if err != nil {
		log.Fatal(err)
	}

	var out bytes.Buffer
	json.Indent(&out, bt, "=", "    ")
	out.WriteTo(os.Stdout)
}
/*
outout:
[
=    {
=        "Name": "Diamond Fork",
=        "Number": 29
=    },
=    {
=        "Name": "Sheep Creek",
=        "Number": 51
=    }
=]%
/*

