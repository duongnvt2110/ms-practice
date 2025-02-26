package config

import (
	"log"
	sharedCfg "pkg/config"
	"sync"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

var (
	cfgOnce sync.Once
	cfg     *Config
)

type Config struct {
	sharedCfg.App
	sharedCfg.Google
	sharedCfg.Mysql
}

func NewConfig() *Config {
	cfgOnce.Do(func() {
		cfg = loadConfig()
	})
	return cfg
}

func loadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file %s", err)
	}
	cfg = &Config{}
	err = envconfig.Process("", cfg)
	if err != nil {
		log.Println("Error loading .env file %s", err)
	}
	return cfg
}
