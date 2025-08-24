//go:build integration

package mongo

import (
	"context"
	"testing"

	"github.com/social-media-monitoring-x/internal/env"
)

var envs env.Enviroments

func TestMain(m *testing.M) {
	envs.Mongo.Host = "localhost"
	envs.Mongo.Port = "27017"
	envs.Mongo.Password = "example"
	envs.Mongo.User = "root"
}

func TestMongoConnection(t *testing.T) {
	t.Parallel()
	t.Run("should connect to mongo, when mongo server is ok", func(t *testing.T) {
		t.Parallel()
		client, err := NewServer(&envs)
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
		auxEnv := envs
		auxEnv.Mongo.Host = "test"
		t.Parallel()
		client, err := NewServer(&auxEnv)
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
