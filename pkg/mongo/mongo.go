package mongo

import (
	"net/url"
	"strings"

	"github.com/social-media-monitoring-x/internal/config"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func NewServer(cfg *config.MongoConfig) (*mongo.Client, error) {
	uri, err := buildMongoURI(cfg)
	if err != nil {
		return nil, err
	}

	return mongo.Connect(
		options.Client().
			ApplyURI(uri).
			SetServerAPIOptions(options.ServerAPI(options.ServerAPIVersion1)),
	)
}

func buildMongoURI(cfg *config.MongoConfig) (string, error) {
	if cfg.Host == "" || cfg.Port == "" {
		return "", NewErrorMongo(nil, "host and port are required")
	}

	encodedUser := url.QueryEscape(cfg.User)
	encodedPassword := url.QueryEscape(cfg.Password)

	uri := strings.Builder{}
	if encodedUser == "" || encodedPassword == "" {
		return "", NewErrorMongo(nil, "user and password are required")
	}

	uri.WriteString("mongodb://")
	uri.WriteString(encodedUser)
	uri.WriteString(":")
	uri.WriteString(encodedPassword)
	uri.WriteString("@")
	uri.WriteString(cfg.Host)
	uri.WriteString(":")
	uri.WriteString(cfg.Port)
	// uri.WriteString("/?authSource=admin")

	return uri.String(), nil
}
