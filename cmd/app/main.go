package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/social-media-monitoring-x/internal/config"
	"github.com/social-media-monitoring-x/internal/container"
	"github.com/social-media-monitoring-x/pkg/fiber"
	"github.com/social-media-monitoring-x/pkg/mongo"
	"go.uber.org/dig"
)

var shutdownChan = make(chan struct{})

func main() {
	appContainer := container.NewContainer()
	registerConfig(appContainer)
	app := fiber.Unwrap(appContainer)
	app.Get("/test", fiber.Test)
	log.Fatal(
		app.Listen(
			config.GetAppAddress(appContainer),
		),
	)
	shutdowngracefully()
	<-shutdownChan
}

// registerConfig use Provide func to registrate a new initialization.
func registerConfig(container *dig.Container) {
	if err := container.Provide(config.LoadAppConfig); err != nil {
		log.Fatalf("failed to provide app config: %v", err)
	}
	if err := container.Provide(config.LoadMongoConfig); err != nil {
		log.Fatalf("failed to provide mongo config: %v", err)
	}
	if err := container.Provide(mongo.NewServer); err != nil {
		log.Fatalf("failed to provide new server to mongo: %v", err)
	}
	if err := container.Provide(fiber.NewServer); err != nil {
		log.Fatalf("failed to provide new server to fiber: %v", err)
	}
}

// shutdowngracefully turn off app gracefully.
func shutdowngracefully() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		fmt.Println("shutting down application gracefully")
		shutdownChan <- struct{}{}
	}()
}
