package auth

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type GoogleConfig struct {
	OAuth *oauth2.Config
}

func NewGoogleConfig(clientID string, clientSecret string, redirectURL string) *GoogleConfig {
	return &GoogleConfig{
		OAuth: &oauth2.Config{
			ClientID:     clientID,
			ClientSecret: clientSecret,
			RedirectURL:  redirectURL,
			Scopes: []string{
				"openid",
				"email",
				"profile",
			},
			Endpoint: google.Endpoint,
		},
	}
}
