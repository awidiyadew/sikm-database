package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type DBCredential struct {
	Host         string
	Username     string
	Password     string
	DatabaseName string
	Port         int
}

func Connect(creds DBCredential) (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta", creds.Host, creds.Username, creds.Password, creds.DatabaseName, creds.Port)

	dbConn, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	return dbConn, nil
}
