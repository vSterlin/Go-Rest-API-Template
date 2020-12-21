package main

import (
	"flag"
	"log"

	"github.com/vSterlin/api-template/internal/config"

	"github.com/BurntSushi/toml"
	"github.com/vSterlin/api-template/internal/apiserver"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := config.Defaults()
	_, err := toml.DecodeFile(configPath, config)

	if err != nil {
		log.Fatal(err)
	}

	if err := apiserver.Start(config); err != nil {
		log.Fatal(err)
	}
}
