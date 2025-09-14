package config

import (
	"fmt"

	"github.com/caarlos0/env"
	"go.uber.org/dig"
)

type AppConfig struct {
	Host string `env:"APP__HOST"`
	Port string `env:"APP__PORT"`
	Name string `env:"APP__NAME"`
}
type MongoConfig struct {
	User     string `env:"MONGO__USER"`
	Password string `env:"MONGO__PASSWORD"`
	Host     string `env:"MONGO__HOST"`
	Port     string `env:"MONGO__PORT"`
}

func LoadAppConfig() (*AppConfig, error) {
	var cfg AppConfig
	if err := env.Parse(&cfg); err != nil {
		// criar um erro e utilizar nesse ponto
		return nil, fmt.Errorf("failed to load app config: %w", err)
	}
	return &cfg, nil
}

func LoadMongoConfig() (*MongoConfig, error) {
	var cfg MongoConfig
	if err := env.Parse(&cfg); err != nil {
		// criar um erro e utilizar nesse ponto
		return nil, fmt.Errorf("failed to load mongo config: %w", err)
	}
	return &cfg, nil
}

func GetAppAddress(d *dig.Container) string {
	var result string
	d.Invoke(func(cfg *AppConfig) {
		result = fmt.Sprintf("%s", cfg.Host+":"+cfg.Port)
	})
	return result
}
