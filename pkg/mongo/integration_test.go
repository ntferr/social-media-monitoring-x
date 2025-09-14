//go:build integration

package mongo

import (
	"context"
	"testing"

	"github.com/social-media-monitoring-x/internal/config"
)

var cfg config.MongoConfig

func TestMain(m *testing.M) {
	cfg.Host = "localhost"
	cfg.Port = "27017"
	cfg.Password = "example"
	cfg.User = "root"
}

func TestMongoConnection(t *testing.T) {
	t.Parallel()
	t.Run("should connect to mongo, when mongo server is ok", func(t *testing.T) {
		t.Parallel()
		client, err := NewServer(&cfg)
		ctx := context.Background()
		if err != nil {
			t.Fatalf("failed to connect to mongo: %v", err)
		}
		defer client.Disconnect(ctx)
		if client == nil {
			t.Fatalf("failed to connect to mongo")
		}
		err = client.Ping(ctx, nil)
		if err != nil {
			t.Fatalf("failed to ping mongo: %v", err)
		}
	})
	t.Run("should return an error, when mongo server is not ok", func(t *testing.T) {
		auxCfg := cfg
		auxCfg.Host = "test"
		t.Parallel()
		client, err := NewServer(&auxCfg)
		ctx := context.Background()
		if err != nil {
			t.Fatalf("failed to connect to mongo: %v", err)
		}
		defer client.Disconnect(ctx)
		if client == nil {
			t.Fatalf("failed to connect to mongo")
		}
		err = client.Ping(ctx, nil)
		if err == nil {
			t.Fatal("ping failed")
		}
	})
}
