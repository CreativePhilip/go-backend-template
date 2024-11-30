package db

import (
	"context"
	"database/sql"
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

type CommonDbTx interface {
	BindNamed(query string, arg interface{}) (string, []interface{}, error)
	DriverName() string
	Get(dest interface{}, query string, args ...interface{}) error
	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	MustExec(query string, args ...interface{}) sql.Result
	MustExecContext(ctx context.Context, query string, args ...interface{}) sql.Result
	NamedExec(query string, arg interface{}) (sql.Result, error)
	NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error)
	NamedQuery(query string, arg interface{}) (*sqlx.Rows, error)
	PrepareNamed(query string) (*sqlx.NamedStmt, error)
	PrepareNamedContext(ctx context.Context, query string) (*sqlx.NamedStmt, error)
	Preparex(query string) (*sqlx.Stmt, error)
	PreparexContext(ctx context.Context, query string) (*sqlx.Stmt, error)
	QueryRowx(query string, args ...interface{}) *sqlx.Row
	QueryRowxContext(ctx context.Context, query string, args ...interface{}) *sqlx.Row
	Queryx(query string, args ...interface{}) (*sqlx.Rows, error)
	QueryxContext(ctx context.Context, query string, args ...interface{}) (*sqlx.Rows, error)
	Rebind(query string) string
	Select(dest interface{}, query string, args ...interface{}) error
	SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
}

var connectionPool *sqlx.DB

func Client() *sqlx.DB {
	if connectionPool != nil {
		return connectionPool
	}

	envCfg := config.GetConfig()
	pool := ClientFromConfig(ClientConfig{
		Host:     envCfg.PostgresHost,
		User:     envCfg.PostgresUser,
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
