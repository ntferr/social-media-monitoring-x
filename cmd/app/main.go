package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
	"github.com/social-media-monitoring-x/internal/config"
	"github.com/social-media-monitoring-x/internal/router"
	pkgFiber "github.com/social-media-monitoring-x/pkg/fiber"
	"github.com/social-media-monitoring-x/pkg/mongo"
	"go.uber.org/dig"
)

var shutdownChan = make(chan struct{})

func main() {
	container := dig.New()
	registerConfig(container)
	if err := container.Invoke(initApp); err != nil {
		log.Fatalf("failed to invoke setup router")
	}
	shutdowngracefully()
	<-shutdownChan
}

// registerConfig use Provide func to registrate a new initialization.
func registerConfig(container *dig.Container) {
	// if err := container.Provide(otel.NewOTLPExporter(ctx)); err != nil {
	// 	log.Fatalf("failed to provide open telemetry: %v", err)
	// }
	if err := container.Provide(config.LoadAppConfig); err != nil {
		log.Fatalf("failed to provide app config: %v", err)
	}
	if err := container.Provide(config.LoadMongoConfig); err != nil {
		log.Fatalf("failed to provide mongo config: %v", err)
	}
	if err := container.Provide(config.LoadXCredentials); err != nil {
		log.Fatalf("failed to provide x credentials: %v", err)
	}
	if err := container.Provide(mongo.NewServer); err != nil {
		log.Fatalf("failed to provide new server to mongo: %v", err)
	}
	if err := container.Provide(pkgFiber.NewServer); err != nil {
		log.Fatalf("failed to provide new server to fiber: %v", err)
	}
}

func initApp(app *fiber.App, cfg *config.AppConfig) {
	router.NewRouter(app).SetupRouter()
	addr := fmt.Sprintf(":%s", cfg.Port)
	if err := app.Listen(addr); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}

// shutdowngracefully shutdown app gracefully.
func shutdowngracefully() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		fmt.Println("shutting down application gracefully")
		shutdownChan <- struct{}{}
	}()
}
