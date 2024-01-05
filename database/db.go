package database

import "database/sql"

type DB interface {
	GetDB() *sql.DB
}
