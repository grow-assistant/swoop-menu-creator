package config

import (
	"fmt"
	"os"
)

const (
	dbHost     = "DB_HOST"
	dbPort     = "DB_PORT"
	dbUser     = "DB_USER"
	dbPassword = "DB_PASSWORD"
	dbName     = "DB_NAME"
)

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

func (c DBConfig) ToConnectionString() string {
	return fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=disable", c.Host, c.Port, c.User, c.Name, c.Password)
}

func NewDBConfig() DBConfig {

	host := getDBHost()
	port := getDBPort()
	name := getDBName()
	user := getDBUser()
	password := getDBPassword()

	c := DBConfig{
		Host:     host,
		Port:     port,
		Name:     name,
		User:     user,
		Password: password,
	}

	return c
}

func getDBHost() string {
	return os.Getenv(dbHost)
}

func getDBPort() string {
	return os.Getenv(dbPort)
}

func getDBName() string {
	return os.Getenv(dbName)
}

func getDBUser() string {
	return os.Getenv(dbUser)
}

func getDBPassword() string {
	return os.Getenv(dbPassword)
}
