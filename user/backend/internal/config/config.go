package config

import (
	"fmt"
	"os"
)

type Config struct {
	Port string
	dbUser string
	dbPass string
	dbHost string
	dbPort string
	dbName string
}

func Dev() Config {
	return Config{
		Port: os.Getenv("port"),
		dbUser: os.Getenv("db-user"),
		dbPass: os.Getenv("db-pass"),
		dbHost: os.Getenv("db-host"),
		dbPort: os.Getenv("db-port"),
		dbName: os.Getenv("db-name"),
	}
}

func (c *Config) DbUrl() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", c.dbUser, c.dbPass,
		c.dbHost, c.dbPort, c.dbName)		
}
