package env

import (
	"github.com/caarlos0/env"
)

type Enviroments struct {
	App struct {
		Host string `env:"APP__HOST"`
		Port string `env:"APP__PORT"`
		Name string `env:"APP__NAME"`
	}
	Mongo struct {
		User     string `env:"MONGO__USER"`
		Password string `env:"MONGO__PASSWORD"`
		Host     string `env:"MONGO__HOST"`
		Port     string `env:"MONGO__PORT"`
	}
}

func loadEnvs() (*Enviroments, error) {
	var enviroments Enviroments
	if err := env.Parse(&enviroments); err != nil {
		return nil, err
	}
	return &enviroments, nil
}
