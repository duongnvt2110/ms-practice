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
		Host string `envconfig:"GRPC_USER_SVC_HOST" default:"user-service"`
		Port string `envconfig:"GRPC_USER_SVC_PORT" default:"50001"`
	}
	GrpcAuthSvc struct {
		Host string `envconfig:"GRPC_AUTH_SVC_HOST" default:"auth-service"`
		Port string `envconfig:"GRPC_AUTH_SVC_PORT" default:"50002"`
	}
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
		log.Printf("Error loading .env file %v", err)
	}
	cfg = &Config{}
	err = envconfig.Process("", cfg)
	if err != nil {
		log.Printf("Error loading .env file %v", err)
	}
	return cfg
}
