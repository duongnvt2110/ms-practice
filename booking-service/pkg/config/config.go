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
	GrpcUserSvc struct {
		UserHost string `envconfig:"GRPC_USER_SVC_HOST" default:"user-service"`
		UserPort string `envconfig:"GRPC_USER_SVC_PORT" default:"50001"`
	}
}

func NewConfig() *Config {
	cfgOnce.Do(func() {
		cfg = LoadConfig()
	})
	return cfg
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file %v", err)
	}
	cfg = &Config{}
	err = envconfig.Process("", cfg)
	if err != nil {
		log.Printf("Error loading .env file %v", err)
	}
	return cfg
}
