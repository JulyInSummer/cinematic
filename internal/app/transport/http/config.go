package http

import (
	"fmt"
	"go.uber.org/config"
	"time"
)

type Config struct {
	Host        string        `yaml:"host"`
	Port        string        `yaml:"port"`
	ReadTimeout time.Duration `yaml:"read_timeout"`
	Mode        string        `yaml:"mode"`
}

func NewConfig(provider config.Provider) (*Config, error) {
	var conf Config
	if err := provider.Get("http").Populate(&conf); err != nil {
		return nil, fmt.Errorf("new config: %w", err)
	}

	return &conf, nil
}
