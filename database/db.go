package database

import "database/sql"

type config struct {
	User     string
	Password string
	Host     string
	Port     int
	DBName   string
	SslMode  string
}

type DB interface {
	GetDB() *sql.DB
}
