package config

import (
	"os"
)

const (
	googleClientID     = "GOOGLE_CLIENT_ID"
	googleClientSecret = "GOOGLE_CLIENT_SECRET"
	googleProjectID    = "GOOGLE_PROJECT_ID"
	googlePubSubCreds  = "GOOGLE_PUBSUB_ENCODED_CREDS"
)

// GoogleConfig represents configuration needed to access Google API
type GoogleConfig struct {
	ClientID          string
	ClientSecret      string
	ProjectID         string
	PubSubCredentials string
}

// NewGoogleConfig initializes new instance of GoogleConfig
func NewGoogleConfig() GoogleConfig {
	return GoogleConfig{
		ClientID:          getGoogleClientID(),
		ClientSecret:      getGoogleClientSecret(),
		ProjectID:         getGoogleProjectID(),
		PubSubCredentials: getGooglePubSubCredentials(),
	}
}

func getGoogleClientID() string {
	return os.Getenv(googleClientID)
}

func getGoogleClientSecret() string {
	return os.Getenv(googleClientSecret)
}

func getGoogleProjectID() string {
	return os.Getenv(googleProjectID)
}

func getGooglePubSubCredentials() string {
	return os.Getenv(googlePubSubCreds)
}
