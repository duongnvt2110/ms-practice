package config

import (
	"log"
	"sync"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

var (
	cfgOnce sync.Once
	cfg     *Config
)

type Config struct {
	Kafka struct {
		Brokers []string `envconfig:"KAFKA_BROKERS" default:"localhost:9092"`
	}
}

func NewConfig() *Config {
	cfgOnce.Do(func() {
		cfg = loadConfig()
	})
	return cfg
}

func loadConfig() *Config {
	_ = godotenv.Load()
	cfg = &Config{}
	if err := envconfig.Process("", cfg); err != nil {
		log.Println("error loading env", err)
	}
	return cfg
}
