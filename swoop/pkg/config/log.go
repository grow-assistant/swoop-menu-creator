package config

import "os"

const (
	logLevel = "LOG_LEVEL"
)

type LogConfig struct {
	LogLevel string
}

func NewLogConfig() LogConfig {
	return LogConfig{
		LogLevel: getLogLevel(),
	}
}

func getLogLevel() string {
	return os.Getenv(logLevel)
}
