package config

import (
	"os"
)

type AppConfig struct {
	PostgreSql *PostgreSql
}

type PostgreSql struct {
	Host     string
	Port     int
	User     string
	Password string
	DbName   string
}

func NewAppConfig() *AppConfig {
	return &AppConfig{
		&PostgreSql{
			Host:     "localhost",
			Port:     5432,
			User:     "postgres",
			Password: os.Getenv("dbPass"),
			DbName:   "geekcontas",
		},
	}
}
