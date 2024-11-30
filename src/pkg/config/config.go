package config

import (
	"fmt"
	cfgo "github.com/CreativePhilip/cfgo/src"
	"os"
)

type AppConfig struct {
	PostgresHost     string `env:"POSTGRES_HOST"`
	PostgresUser     string `env:"POSTGRES_USER"`
	PostgresPassword string `env:"POSTGRES_PASSWORD"`
	PostgresDatabase string `env:"POSTGRES_DB"`

	DBSchemaLocation string `env:"SCHEMA_LOCATION"`

	PasswordSalt string `env:"PASSWORD_SALT"`
}

func GetConfig() AppConfig {
	appConfig := AppConfig{}

	cfgo.LoadType(&appConfig, cfgo.NewEnvConfiguration(cfgo.EnvConfiguration{
		BoolValidTrueValues: nil,
		Providers: []cfgo.ConfigSourceProvider{
			cfgo.NewEnvFileVariableSourceProvider(fmt.Sprintf("%s/env/local/*.env", os.Getenv("ENV_FILES_DIR"))),
			cfgo.NewEnvVariablesSourceProvider(),
		},
	}))

	return appConfig
}
