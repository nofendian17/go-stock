package config

type Application struct {
	Name     string `mapstructure:"name"`
	Version  string `mapstructure:"version"`
	Timezone string `mapstructure:"timezone"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
}
