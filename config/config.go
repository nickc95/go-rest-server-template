package config

import (
	"log"
	"sync"

	"github.com/caarlos0/env/v6"
)

type ConfigObject struct {
	ApplicationName string `env:"ApplicationName" envDefault:"go-rest-server-template"`
	Port string `env:"Port" envDefault:"3000"`
	Version string `env:"Version" envDefault:"0.1.0"`
}

var instantiated *ConfigObject

var once sync.Once

func read() *ConfigObject {
	config := ConfigObject{}
	if err := env.Parse(&config); err != nil {
		log.Fatal(err)
	}
	return &config
}

func GetInstance() *ConfigObject {
	once.Do(func() {
		instantiated = read()
	})
	return instantiated
}
