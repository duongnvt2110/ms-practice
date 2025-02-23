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
	App struct {
		Host string `envconfig:"APP_HOST" default:"localhost"`
		Port string `envconfig:"APP_PORT" default:"3000"`
	}

	Google struct {
		OauthClientID     string   `envconfig:"GOOGLE_OAUTH_CLIENT_ID"`
		OauthClientSecret string   `envconfig:"GOOGLE_OAUTH_CLIENT_SECRET"`
		OauthScopes       []string `envconfig:"GOOGLE_OAUTH_SCOPES" default:"https://www.googleapis.com/auth/userinfo.email"`
		OauthGoogleUrlAPI string   `envconfig:"GOOGLE_OAUTH_URL_API" default:"https://www.googleapis.com/oauth2/v2/userinfo?access_token="`
	}

	Mysql struct {
		Host     string `envconfig:"MYSQL_HOST" default:"localhost"`
		Port     string `envconfig:"MYSQL_PORT" default:"3366"`
		User     string `envconfig:"MYSQL_USER" default:"admin"`
		Password string `envconfig:"MYSQL_PASSWORD" default:"admin"`
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
		log.Println("Error loading .env file %s", err)
	}
	cfg = &Config{}
	err = envconfig.Process("", cfg)
	if err != nil {
		log.Println("Error loading .env file %s", err)
	}
	return cfg
}
