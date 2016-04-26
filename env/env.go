package env

import (
	"os"
)

var v variables

type variables struct {
	DB Database
}

// Database ...
type Database struct {
	Name     string
	Host     string
	Port     string
	User     string
	Password string
}

// GetDB return database configuration
func GetDB() Database {
	return v.DB
}

// Load ...
func Load() {
	v = variables{
		DB: Database{
			Name:     os.Getenv("db_name"),
			Host:     os.Getenv("db_host"),
			Port:     os.Getenv("db_port"),
			User:     os.Getenv("db_user"),
			Password: os.Getenv("db_password"),
		},
	}
}
