package config

import (
	"log"
	"sync"

	sharedcfg "ms-practice/pkg/config"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

var (
	cfgOnce sync.Once
	cfg     *Config
)

type Config struct {
	sharedcfg.App
	Notification Notification `envconfig:""`
	SMTP         SMTP
	Firebase     Firebase
}

type Notification struct {
	MaxAttempts          int `envconfig:"NOTI_MAX_ATTEMPTS" default:"3"`
	RetryIntervalSeconds int `envconfig:"NOTI_RETRY_INTERVAL_SECONDS" default:"60"`
	WorkerIntervalMS     int `envconfig:"NOTI_WORKER_INTERVAL_MS" default:"500"`
}

type SMTP struct {
	Host     string `envconfig:"SMTP_HOST" default:"localhost"`
	Port     int    `envconfig:"SMTP_PORT" default:"1025"`
	Username string `envconfig:"SMTP_USERNAME"`
	Password string `envconfig:"SMTP_PASSWORD"`
	From     string `envconfig:"SMTP_FROM" default:"noreply@example.com"`
}

type Firebase struct {
	CredentialsFile string `envconfig:"FIREBASE_CREDENTIALS_FILE" default:""`
	ProjectID       string `envconfig:"FIREBASE_PROJECT_ID" default:""`
}

func NewConfig() *Config {
	cfgOnce.Do(func() {
		cfg = loadConfig()
	})
	return cfg
}

func loadConfig() *Config {
	_ = godotenv.Load()
	c := &Config{}
	if err := envconfig.Process("", c); err != nil {
		log.Printf("failed to load notification config: %v", err)
	}
	if c.Notification.MaxAttempts <= 0 {
		c.Notification.MaxAttempts = 3
	}
	if c.Notification.RetryIntervalSeconds <= 0 {
		c.Notification.RetryIntervalSeconds = 60
	}
	if c.Notification.WorkerIntervalMS <= 0 {
		c.Notification.WorkerIntervalMS = 500
	}
	return c
}
