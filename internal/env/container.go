package env

import (
	"log"

	"go.uber.org/dig"
)

var container *dig.Container

func init() {
	container = dig.New()
	if err := container.Provide(loadEnvs); err != nil {
		log.Fatalf("failed to provide envs: %v", err)
	}
}

func Resolve(fn interface{}) error {
	return container.Invoke(fn)
}
