package env

import (
	"log"

	"go.uber.org/dig"
)

var Container *dig.Container

func init() {
	Container = dig.New()
	if err := Container.Provide(loadEnvs); err != nil {
		log.Fatalf("failed to provide envs: %v", err)
	}
}

func Resolve(fn interface{}) error {
	return Container.Invoke(fn)
}
