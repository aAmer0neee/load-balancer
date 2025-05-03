package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

var (
	configPath = flag.String("config", "", "path to configure file")
)

func MustLoad(cfg interface{}) {

	if err := cleanenv.ReadConfig(cfgPath(), cfg); err != nil {
		log.Fatalf("error load config %v", err)
	}

}

func cfgPath() string {

	flag.Parse()

	if *configPath == "" {
		log.Fatalf("path to configure file not specified")
	}

	if _, err := os.Stat(*configPath); err == os.ErrNotExist {
		log.Fatalf("no such file %s", *configPath)
	}

	return *configPath
}
