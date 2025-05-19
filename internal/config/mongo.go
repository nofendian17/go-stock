package config

type Mongo struct {
	DSN      string `mapstructure:"dsn"`
	Database string `mapstructure:"database"`
}
