package config

import (
	"os"
)

const (
	teamID   = "APPLE_TEAM_ID"
	clientID = "APPLE_CLIENT_ID"
	keyID    = "APPLE_KEY_ID"
	secret   = "APPLE_SECRET"
)

// AppleConfig represents configuration needed to access Apple API
type AppleConfig struct {
	TeamID   string
	ClientID string
	KeyID    string
	Secret   string
}

// NewAppleConfig initializes new instnce of AppleConfig
func NewAppleConfig() AppleConfig {
	return AppleConfig{
		TeamID:   getAppleTeamID(),
		ClientID: getAppleClientID(),
		KeyID:    getAppleKeyID(),
		Secret:   getAppleSecret(),
	}
}

func getAppleTeamID() string {
	return os.Getenv(teamID)
}

func getAppleClientID() string {
	return os.Getenv(clientID)
}

func getAppleKeyID() string {
	return os.Getenv(keyID)
}

func getAppleSecret() string {
	return os.Getenv(secret)
}
