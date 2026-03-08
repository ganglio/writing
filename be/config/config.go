package config

import (
	"encoding/json"
	"errors"
	"os"

	"dario.cat/mergo"
	"github.com/caarlos0/env/v6"
	log "github.com/sirupsen/logrus"
)

type Env struct {
	AppEnv     string `env:"APP_ENV" envDefault:"production" json:"-"`
	Port       uint16 `env:"PORT" envDefault:"3000" json:"port"`
	WorkFolder string `env:"WORK_FOLDER" envDefault:"." json:"-"`

	FEDevHost string `env:"FE_DEV_HOST" envDefault:"localhost" json:"-"`
	FEDevPort uint16 `env:"FE_DEV_PORT" envDefault:"3001" json:"-"`

	Tabs []Tab `json:"tabs"`
}

func (e *Env) IsDev() bool {
	return e.AppEnv != "production"
}

var envConfig *Env

func GetEnv() *Env {
	var (
		cfg  Env
		cfg2 Env
	)

	// Memoize FTW!
	if envConfig != nil {
		return envConfig
	}

	// Load environment variables into cfg
	if err := env.Parse(&cfg); err != nil {
		log.WithError(err).Fatal("Failed to parse environment variables, using defaults")
	}

	// Prepopulate default tabs
	if len(cfg.Tabs) == 0 {
		cfg.Tabs = DefaultTabs
	}

	// Load .writing.json if it exists and merge it with environment variables
	// If .writing.json doesn't exist, create it with the current environment variables
	jsonConfig, err := os.ReadFile(cfg.WorkFolder + "/.writing.json")
	if errors.Is(err, os.ErrNotExist) {
		log.Warn("No .writing.json found, populating it from environment variables.")
		err := cfg.Save()
		if err != nil {
			log.WithError(err).Fatal("Failed to save .writing.json.")
		}
		goto skip
	} else if err != nil {
		log.WithError(err).Fatal("Failed to read .writing.json.")
	}

	if err := json.Unmarshal(jsonConfig, &cfg2); err != nil {
		log.WithError(err).Fatal("Failed to unmarshal .writing.json.")
	}

	if err := mergo.Merge(&cfg, cfg2); err != nil {
		log.WithError(err).Fatal("Failed to merge .writing.json with environment variables.")
	}

skip:
	envConfig = &cfg
	return envConfig
}

func (e *Env) Save() error {
	out, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(e.WorkFolder+"/.writing.json", out, 0644)
}

func (e *Env) AddTab() string {
	return "pescolone"
}
