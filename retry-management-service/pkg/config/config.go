package config

import (
	"log"
	"sync"

	sharedCfg "ms-practice/pkg/config"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

var (
	cfgOnce sync.Once
	cfg     *Config
)

type Config struct {
	sharedCfg.App
	DlqTopic   string `envconfig:"DLQ_TOPIC" default:"dlq.events"`
	DlqGroupID string `envconfig:"DLQ_GROUP_ID" default:"retry-management-service"`
}

func NewConfig() *Config {
	cfgOnce.Do(func() {
		cfg = loadConfig()
	})
	return cfg
}

func loadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Printf("Error loading .env file %v", err)
	}
	cfg = &Config{}
	if err := envconfig.Process("", cfg); err != nil {
		log.Printf("Error loading env config %v", err)
	}
	return cfg
}
