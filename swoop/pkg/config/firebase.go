package config

import (
	"os"
)

const (
	firebaseEncodedCreds = "FIREBASE_ENCODED_CREDS"
)

// FirebaseConfig represents configuration needed to access Firebase API
type FirebaseConfig struct {
	Credentials string
}

// NewFirebaseConfig initializes new instance of FirebaseConfig
func NewFirebaseConfig() FirebaseConfig {
	return FirebaseConfig{
		Credentials: getFirebaseEncodedCreds(),
	}
}

func getFirebaseEncodedCreds() string {
	return os.Getenv(firebaseEncodedCreds)
}
