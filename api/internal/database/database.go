package database

import (
	"database/sql"
	"log"

	"github.com/felipeazsantos/personal-finance-go/config"
	_ "github.com/lib/pq"
)

func New(appConfig *config.AppConfig) *sql.DB {
	conn, err := sql.Open("postgres", appConfig.ConnStr)
	if err != nil {
		log.Fatalf("cannot connect to database: %v", err)
	}

	if err := conn.Ping(); err != nil {
		log.Fatalf("database is not responding: %v", err)
	}

	return conn
}
