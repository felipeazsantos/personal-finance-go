package config

import "fmt"

const (
	user     = "felipe"
	password = "123"
	dbName   = "personal_finance"
	host     = "localhost"
	port     = 5433
	dsn      = "host=%s port=%d user=%s password=%s dbname=%s sslmode=disable"
)

type AppConfig struct {
	ConnStr string
}

func New() *AppConfig {
	return &AppConfig{
		ConnStr: fmt.Sprintf(dsn, host, port, user, password, dbName),
	}
}
