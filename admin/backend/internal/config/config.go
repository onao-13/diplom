package config

import (
	"fmt"
	"os"
)

type Config struct {
	Port   string
	dbHost string
	dbPort string
	dbUser string
	dbPass string
	dbName string
	Root   RootUser
}

func Load() Config {
	return Config{
		Port:   os.Getenv("port"),
		dbHost: os.Getenv("db-host"),
		dbPort: os.Getenv("db-port"),
		dbUser: os.Getenv("db-user"),
		dbPass: os.Getenv("db-pass"),
		dbName: os.Getenv("db-name"),
		Root: RootUser{
			Username: os.Getenv("root-username"),
			Password: os.Getenv("root-password"),
		},
	}
}

func (c *Config) DbUrl() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", c.dbUser, c.dbPass,
		c.dbHost, c.dbPort, c.dbName)
}
