package postgres

import (
	"fmt"
	"go.uber.org/config"
)

type Config struct {
	PostgresPassword string `yaml:"postgres_password"`
	PostgresUser     string `yaml:"postgres_user"`
	PostgresHost     string `yaml:"postgres_host"`
	PostgresPort     int    `yaml:"postgres_port"`
	PostgresDatabase string `yaml:"postgres_database"`
}

func NewConfig(provider config.Provider) (*Config, error) {
	var conf Config

	err := provider.Get("postgres").Populate(&conf)
	if err != nil {
		return nil, fmt.Errorf("unable to parse postgres config: %w", err)
	}

	return &conf, nil
}
