package config

import (
	"os"
)

const (
	facebookClientID     = "FACEBOOK_CLIENT_ID"
	facebookClientSecret = "FACEBOOK_CLIENT_SECRET"
)

// FacebookConfig represents configuration needed to access Facebook API
type FacebookConfig struct {
	ClientID     string
	ClientSecret string
}

// NewFacebookConfig initializes new instnce of FacebookConfig
func NewFacebookConfig() FacebookConfig {
	return FacebookConfig{
		ClientID:     getClientID(),
		ClientSecret: getClientSecret(),
	}
}

func getClientID() string {
	return os.Getenv(facebookClientID)
}

func getClientSecret() string {
	return os.Getenv(facebookClientSecret)
}
