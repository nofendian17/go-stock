package config

import (
	"github.com/spf13/viper"
)

type Config interface {
	GetApplication() Application
	GetMongo() Mongo
	GetService() Service
	GetCronJob() CronJob
}

type config struct {
	Application Application `mapstructure:"application"`
	Mongo       Mongo       `mapstructure:"mongo"`
	Service     Service     `mapstructure:"service"`
	CronJob     CronJob     `mapstructure:"cron_job"`
}

func (c *config) GetApplication() Application { return c.Application }
func (c *config) GetMongo() Mongo             { return c.Mongo }
func (c *config) GetService() Service         { return c.Service }
func (c *config) GetCronJob() CronJob {
	return c.CronJob
}

func NewConfig(path string) (Config, error) {
	v := viper.New()
	v.SetConfigFile(path)
	v.SetConfigType("yaml") // or "json", "toml" depending on your config file

	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	cfg := &config{}
	if err := v.Unmarshal(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
