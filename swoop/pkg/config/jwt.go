package config

import (
	"os"
)

const (
	jwtSigningKeyEnv = "JWT_SECRET"
)

// JWTConfig represents configuration needed for JWT operations
type JWTConfig struct {
	SigningKey string
}

// NewJWTConfig initializes new instance of JWTConfig
func NewJWTConfig() JWTConfig {
	return JWTConfig{
		SigningKey: getJWTSigningKey(),
	}
}

func getJWTSigningKey() string {
	return os.Getenv(jwtSigningKeyEnv)
}
