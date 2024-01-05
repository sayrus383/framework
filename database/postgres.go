package database

import (
	"database/sql"
	"fmt"
	"gitlab.bcc.kz/digital-banking-platform/microservices/currency-market/modules/framework.git/application"
	"gitlab.bcc.kz/digital-banking-platform/microservices/currency-market/modules/framework.git/errors"
	"gitlab.bcc.kz/digital-banking-platform/microservices/currency-market/modules/framework.git/logger"

	_ "github.com/lib/pq"
)

type postgres struct {
	db *sql.DB
}

func NewPostgresConn(cfg application.Config, log logger.Logger) (DB, error) {
	log = log.WithComponent("database")

	log.Info("database init...")

	dbCfg := config{}

	cfg.StringVar(&dbCfg.DBName, "APP_PG_DBNAME", "postgres", "db name")
	cfg.StringVar(&dbCfg.User, "APP_PG_LOGIN", "postgres", "db user")
	cfg.StringVar(&dbCfg.Password, "APP_PG_PASSWORD", "password", "db password")
	cfg.StringVar(&dbCfg.Host, "APP_PG_HOST", "localhost", "db host")
	cfg.IntVar(&dbCfg.Port, "APP_PG_PORT", 5432, "db port")
	cfg.StringVar(&dbCfg.SslMode, "APP_PG_SSL", "disable", "db support ssl mode")

	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=%s",
		dbCfg.User, dbCfg.Password, dbCfg.Host, dbCfg.Port, dbCfg.DBName, dbCfg.SslMode)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, errors.Wrap("NewPostgresConn sql.Open", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, errors.Wrap("NewPostgresConn db.Ping", err)
	}

	log.Info("database connected")

	return &postgres{
		db: db,
	}, nil
}

func (p postgres) GetDB() *sql.DB {
	return p.db
}
