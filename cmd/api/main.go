package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/LLoylento4bIx4ervei/WebServer/internal/app/api"
)

var (
	configPath = "configs/api.toml"
)

func init() {
	flag.StringVar(&configPath, "path", "configs/api.toml", "path to config file in .toml format")
}

func main() {
	flag.Parse()
	log.Println("hi")
	config := api.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Println("can not fing configs file:", err)
	}
	server := api.New(config)

	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
