package db

import (
	"database/sql"
	"time"

	"github.com/io-m/bankapp/internal/logger"
)

func NewDBClient(driver, dns string) *sql.DB {
	client, err := sql.Open(driver, dns)
	if err != nil {
		logger.Warn.Fatal("Could not CONNECT to DB")
	}
	client.SetConnMaxIdleTime(5 * time.Minute)
	client.SetMaxOpenConns(5)
	client.SetMaxIdleConns(10)
	return client
}