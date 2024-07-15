package config

type App struct {
	AppName string `mapstructure:"name"`
	AppPort string `mapstructure:"port"`
	AppEnv  string `mapstructure:"env"`
}
