package service

import (
	"fmt"
	"go.uber.org/config"
)

type Config struct {
	Secret string `yaml:"secret"`
}

func NewConfig(provider config.Provider) (*Config, error) {
	var conf Config

	err := provider.Get("service").Populate(&conf)
	if err != nil {
		return nil, fmt.Errorf("unable to parse postgres config: %w", err)
	}

	return &conf, nil
}
