package config

import (
	"log"
	"os"

	"github.com/google/uuid"

	"github.com/sirupsen/logrus"

	"github.com/spf13/viper"
)

var instanceID string

func init() {
	u := uuid.New()
	instanceID = u.String()
	logrus.WithField("instanceID", instanceID).Info("initializing application config")
}

const (
	portEnv = "PORT"
)

func Init() error {
	viper.Set(portEnv, getPort())
	return validate()
}

func Port() string {
	return viper.GetString(portEnv)
}

func DB() DBConfig {
	return NewDBConfig()
}

func Facebook() FacebookConfig {
	return NewFacebookConfig()
}

func Google() GoogleConfig {
	return NewGoogleConfig()
}

func Apple() AppleConfig {
	return NewAppleConfig()
}

func JWT() JWTConfig {
	return NewJWTConfig()
}

func Log() LogConfig {
	return NewLogConfig()
}

func Redis() RedisConfig {
	return NewRedisConfig()
}

func InstanceID() string {
	return instanceID
}

func Firebase() FirebaseConfig {
	return NewFirebaseConfig()
}

func getPort() string {
	port := os.Getenv(portEnv)
	if port == "" {
		log.Println("Couldn't find PORT, defaulting to 8080")
		port = "8080"
	}
	return port
}

func validate() error {
	if Port() == "" {
		return NewErrInvalidConfig(portEnv)
	}

	return nil
}
