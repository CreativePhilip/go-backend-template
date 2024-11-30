package utils

import (
	"bytes"
	"database/sql"
	"fmt"
	"github.com/CreativePhilip/backend/src/db"
	"github.com/CreativePhilip/backend/src/pkg/config"
	"github.com/jmoiron/sqlx"
	"os/exec"
	"strings"
)

var hasRanMigrations = false

func SetupIntegrationTest() (*sqlx.DB, func()) {
	return testClient()
}

func testClient() (*sqlx.DB, func()) {
	envCfg := config.GetConfig()
	cfg := db.ClientConfig{
		Host:     envCfg.PostgresHost,
		User:     envCfg.PostgresDb,
		Password: envCfg.PostgresPassword,
		Database: envCfg.PostgresDatabase + "_test",
	}

	if !hasRanMigrations {
		createTestDatabase(cfg)
		runMigrations(envCfg, cfg)

		hasRanMigrations = true
	}

	d := db.ClientFromConfig(cfg)

	return d, func() {
		d.MustExec(`
do $$  declare 
	r RECORD;
begin
for r in (select tablename from pg_tables where schemaname = current_schema()) loop
    execute 'truncate table ' || quote_ident(r.tablename) || ' cascade';
end loop;
end;
$$
`)

		d.MustExec(`
do $$ declare 
	i TEXT;
begin
 for i in (select column_default from information_schema.columns where column_default similar to 'nextval%') 
  LOOP
         EXECUTE 'ALTER SEQUENCE'||' ' || substring(substring(i from '''[a-z_]*')from '[a-z_]+') || ' '||' RESTART 1;';    
  END LOOP;
END $$; 
`)

		d.Close()
	}
}

func createTestDatabase(cfg db.ClientConfig) {
	rawDb := connectToTestInstanceWithoutDb(cfg)

	_, err := rawDb.Exec(fmt.Sprintf("create database %s", cfg.Database))

	if err != nil {
		if !strings.Contains(err.Error(), "already exists") {
			panic(err)
		}
	}

	rawDb.Close()
}

func runMigrations(cfg config.AppConfig, dbConfig db.ClientConfig) {
	cmd := exec.Command(
		"atlas",
		"schema",
		"apply",
		"--url",
		fmt.Sprintf(
			"postgres://%s:%s@%s:5432/%s?sslmode=disable&search_path=public",
			dbConfig.User,
			dbConfig.Password,
			dbConfig.Host,
			dbConfig.Database,
		),
		"--to",
		fmt.Sprintf("file://%s", cfg.DBSchemaLocation),
		"--dev-url",
		"docker://postgres/16",
		"--auto-approve",
	)

	var outb, errb bytes.Buffer
	cmd.Stdout = &outb
	cmd.Stderr = &errb

	err := cmd.Run()

	if err != nil {
		fmt.Println(outb.String(), errb.String())
		fmt.Println(fmt.Sprintf(
			"postgres://%s:%s@%s:5432/%s?sslmode=disable&search_path=public",
			dbConfig.User,
			dbConfig.Password,
			dbConfig.Host,
			dbConfig.Database,
		))
		fmt.Println(fmt.Sprintf("file://%s", cfg.DBSchemaLocation))
		panic(fmt.Errorf("test migration: %w", err))
	}
}

func connectToTestInstanceWithoutDb(cfg db.ClientConfig) *sql.DB {
	rawDb, err := sql.Open(
		"postgres",
		fmt.Sprintf(
			"host=%s user=%s password=%s sslmode=disable",
			cfg.Host,
			cfg.User,
			cfg.Password,
		),
	)
	if err != nil {
		panic(fmt.Errorf("failed to connect to test database: %w", err))
	}

	return rawDb
}
