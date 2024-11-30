package config

import cfgo "github.com/CreativePhilip/cfgo/src"

type AppConfig struct {
	PostgresHost     string `env:"POSTGRES_HOST"`
	PostgresDb       string `env:"POSTGRES_DB"`
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
			cfgo.NewEnvFileVariableSourceProvider("/Users/philip/code/fitsy-booksy-yolo/backend/env/local/*.env"),
		},
	}))

	return appConfig
}
