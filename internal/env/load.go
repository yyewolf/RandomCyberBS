package env

import (
	"rcbs/internal/values"

	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

var cfg config

func Load() {
	godotenv.Load()

	if err := env.Parse(&cfg); err != nil {
		logrus.Fatalf("failed to load env: %v", err)
	}

	if cfg.Mode == values.Unset {
		logrus.Fatalf("MODE is not set, be sure to have a .env file or set the environment variables")
	}

	logrus.SetLevel(logrus.DebugLevel)
	if cfg.Mode == values.Prod {
		logrus.SetLevel(logrus.InfoLevel)
	}

	logrus.Infof("Environment loaded: %s", cfg.Mode)
	logrus.Debugf("Environment: %+v", cfg)
}
