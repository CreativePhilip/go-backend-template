package db

import (
	"fmt"
	"github.com/CreativePhilip/backend/src/pkg/config"
)

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type ClientConfig struct {
	Host     string
	User     string
	Password string
	Database string
}

var connectionPool *sqlx.DB

func Client() *sqlx.DB {
	if connectionPool != nil {
		return connectionPool
	}

	envCfg := config.GetConfig()
	pool := ClientFromConfig(ClientConfig{
		Host:     envCfg.PostgresHost,
		User:     envCfg.PostgresDb,
		Password: envCfg.PostgresPassword,
		Database: envCfg.PostgresDatabase,
	})

	connectionPool = pool
	return pool
}

func ClientFromConfig(cfg ClientConfig) *sqlx.DB {
	dataSource := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host,
		cfg.User,
		cfg.Password,
		cfg.Database,
	)
	db, err := sqlx.Open("postgres", dataSource)

	if err != nil {
		panic(fmt.Sprintf("failed to connect to database: %v", err))
	}

	return db
}
