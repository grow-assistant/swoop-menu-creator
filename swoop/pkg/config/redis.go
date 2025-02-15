package config

import (
	"os"
	"strconv"
)

const (
	redisHostEnv = "REDIS_HOST"
	redisPortEnv = "REDIS_PORT"
)

type RedisConfig struct {
	Host string
	Port int
}

func NewRedisConfig() RedisConfig {

	host := getRedisHost()
	port := getRedisPort()

	return RedisConfig{
		Host: host,
		Port: port,
	}
}

func getRedisHost() string {
	return os.Getenv(redisHostEnv)
}

func getRedisPort() int {
	port, err := strconv.Atoi(os.Getenv(redisPortEnv))
	if err != nil {
		port = 6379
	}
	return port
}
