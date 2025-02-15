package config

import "fmt"

const (
	errorInvalidConfigMsg = "invalid configuration"
	errorInitConfigMsg    = "error initializing configuration"
)


// ErrInvalidConfig occurs when an error occurs during config initialization
type ErrInvalidConfig struct {
	message string
}

// NewErrInvalidConfig returns an instance of ErrorInvalidConfig
func NewErrInvalidConfig(message string) ErrInvalidConfig {
	return ErrInvalidConfig{
		message: message,
	}
}

// Error returns human readable string of ErrInvalidConfig
func (e ErrInvalidConfig) Error() string {
	return fmt.Sprintf("config: %v : %v", errorInvalidConfigMsg, e.message)
}

// ErrInitConfig occurs when an error occurs during config initialization
type ErrInitConfig struct {
	message string
}

// NewErrInitConfig returns an instance of ErrInitConfig
func NewErrInitConfig(message string) ErrInitConfig {
	return ErrInitConfig{
		message: message,
	}
}

// Error returns human readable string of ErrInitConfig
func (e ErrInitConfig) Error() string {
	return fmt.Sprintf("config: %v : %v", errorInitConfigMsg, e.message)
}
