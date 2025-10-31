package config

import (
	"fmt"

	"github.com/caarlos0/env"
	"go.uber.org/dig"
)

// AppConfig struct represents APP environments.
type AppConfig struct {
	Host string `env:"APP__HOST"`
	Port string `env:"APP__PORT"`
	Name string `env:"APP__NAME"`
}

// MongoConfig struct represents Mongo environments.
type MongoConfig struct {
	User     string `env:"MONGO__USER"`
	Password string `env:"MONGO__PASSWORD"`
	Host     string `env:"MONGO__HOST"`
	Port     string `env:"MONGO__PORT"`
}

// XCredentials struct represents x credentials environments.
type XCredentials struct {
	BearerToken string `env:"BEARER_TOKEN"`
}

// LoadAppConfig load app environments to Application.
func LoadAppConfig() (*AppConfig, error) {
	var cfg AppConfig
	if err := env.Parse(&cfg); err != nil {
		// criar um erro e utilizar nesse ponto
		return nil, fmt.Errorf("failed to load app config: %w", err)
	}
	return &cfg, nil
}

// LoadMongoConfig load mongo environments to Application.
func LoadMongoConfig() (*MongoConfig, error) {
	var cfg MongoConfig
	if err := env.Parse(&cfg); err != nil {
		// criar um erro e utilizar nesse ponto
		return nil, fmt.Errorf("failed to load mongo config: %w", err)
	}
	return &cfg, nil
}

// LoadXCredentials load x credentials environments to Application.
func LoadXCredentials() (*XCredentials, error) {
	var credentials XCredentials
	if err := env.Parse(&credentials); err != nil {
		return nil, fmt.Errorf("failed to load x credentials: %w", err)
	}
	return &credentials, nil
}

// GetAppAddress get host and port from environments
// and return address string.
func GetAppAddress(d *dig.Container) string {
	var result string
	d.Invoke(func(cfg *AppConfig) {
		result = cfg.Host + ":" + cfg.Port
	})
	return result
}
