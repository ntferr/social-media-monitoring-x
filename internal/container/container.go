package container

import (
	"go.uber.org/dig"
)

func NewContainer() *dig.Container {
	return dig.New()
}

func Resolve(container *dig.Container, fn interface{}) error {
	return container.Invoke(fn)
}
