package config

import (
	"github.com/caarlos0/env/v6"
)

type Env struct {
	AppEnv     string `env:"APP_ENV" envDefault:"production"`
	Port       uint16 `env:"PORT" envDefault:"3000"`
	Db         string `env:"APP_DB" envDefault:"app.db"`
	WorkFolder string `env:"WORK_FOLDER" envDefault:"."`

	FEDevHost string `env:"FE_DEV_HOST" envDefault:"localhost"`
	FEDevPort uint16 `env:"FE_DEV_PORT" envDefault:"3001"`
}

func (e *Env) IsDev() bool {
	return e.AppEnv != "production"
}

var envConfig *Env

func GetEnv() *Env {
	if envConfig == nil {
		cfg := &Env{}
		if err := env.Parse(cfg); err != nil {
			panic(err)
		}
		envConfig = cfg
	}
	return envConfig
}
