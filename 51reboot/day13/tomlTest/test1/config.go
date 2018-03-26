package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
)

// UserScriptConfig for user defined scrpit config
type UserScriptConfig struct {
	Path string
	Step string
}

type config struct {
	TransAddr  string
	UserScript []UserScriptConfig
}

var (
	configPath = flag.String("config", "config.toml", "config path")
	gcfg       config
)

func main() {
	flag.Parse()

	_, err := toml.DecodeFile(*configPath, &gcfg)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("TransAddr : %#v\n", gcfg.TransAddr)
	for _, script := range gcfg.UserScript {
		log.Printf("UserScript: %#v\n", script)
	}
}
