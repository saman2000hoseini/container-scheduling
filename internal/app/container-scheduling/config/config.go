package config

import (
	"github.com/saman2000hoseini/mossgow/pkg/config"
	"github.com/sirupsen/logrus"
	"gopkg.in/go-playground/validator.v9"
	"time"
)

const (
	app       = "container-scheduling"
	cfgFile   = "config.yaml"
	cfgPrefix = "container-scheduling"
)

type (
	Config struct {
		Server Server `mapstructure:"server"`
	}

	Server struct {
		Address         string        `mapstructure:"address"`
		GracefulTimeout time.Duration `mapstructure:"graceful-timeout"`
		ReadTimeout     time.Duration `mapstructure:"read-timeout"`
		WriteTimeout    time.Duration `mapstructure:"write-timeout"`
	}
)

func (c Config) Validate() error {
	return validator.New().Struct(c)
}

// Init initializes application configuration.
func Init() Config {
	var cfg Config

	config.Init(app, cfgFile, &cfg, defaultConfig, cfgPrefix)

	if err := cfg.Validate(); err != nil {
		logrus.Fatalf("failed to validate configurations: %s", err.Error())
	}

	return cfg
}
