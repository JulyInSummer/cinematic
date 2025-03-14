package app

import (
	"fmt"
	"go.uber.org/config"
	"go.uber.org/fx"
	"os"
)

type Config struct {
	Name string `yaml:"name"`
}

type ResultConfig struct {
	fx.Out
	Provider config.Provider
	Config   Config
}

func NewConfig() (ResultConfig, error) {
	f, err := os.Open("config.local.yaml")
	if err != nil {
		return ResultConfig{}, fmt.Errorf("failed to open config file: %w", err)
	}

	loader, err := config.NewYAML(config.Source(f))
	if err != nil {
		return ResultConfig{}, fmt.Errorf("failed to load config file: %w", err)
	}

	conf := Config{
		Name: "default",
	}

	if err = loader.Get("app").Populate(&conf); err != nil {
		return ResultConfig{}, fmt.Errorf("failed to populate config: %w", err)
	}

	return ResultConfig{
		Provider: loader,
		Config:   conf,
	}, nil
}
