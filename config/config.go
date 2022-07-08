package config

import "github.com/vrischmann/envconfig"

var appconfig Config

type Config struct {
	Database struct {
		URL     string `envconfig:"default=postgres://postgres:mysecretpassword@localhost:5432/learning?sslmode=disable,optional"`
		Logmode bool
	}
	Port       string
	APPVersion string
}

func Initconfig() {
	appconfig = Config{}
	envconfig.Init(&appconfig)
}

func GetConfig() Config {
	return appconfig
}
