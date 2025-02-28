package goauth2

import (
	"ms-practice/auth-service/pkg/config"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func GetOauth2Config(cfg *config.Config, redirectUrl string) *oauth2.Config {
	return &oauth2.Config{
		RedirectURL:  redirectUrl,
		ClientID:     cfg.Google.OauthClientID,
		ClientSecret: cfg.Google.OauthClientSecret,
		Scopes:       cfg.Google.OauthScopes,
		Endpoint:     google.Endpoint,
	}
}
