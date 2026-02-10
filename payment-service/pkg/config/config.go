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
	Kafka struct {
		Brokers []string `envconfig:"KAFKA_BROKERS" default:"localhost:9092"`
	}
	GrpcPaymentSvc struct {
		Host string `envconfig:"GRPC_PAYMENT_SVC_HOST" default:"payment-service"`
		Port string `envconfig:"GRPC_PAYMENT_SVC_PORT" default:"50003"`
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
