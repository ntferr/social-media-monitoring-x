package mongo

import (
	"net/url"
	"strings"

	"github.com/social-media-monitoring-x/internal/env"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func NewServer(envs *env.Enviroments) (*mongo.Client, error) {
	uri, err := buildMongoURI(envs.Mongo.User, envs.Mongo.Password, envs.Mongo.Host, envs.Mongo.Port)
	if err != nil {
		return nil, err
	}
	return mongo.Connect(
		options.Client().ApplyURI(uri),
	)
}

func buildMongoURI(user, password, host, port string) (string, error) {
	if host == "" || port == "" {
		return "", NewErrorMongo(nil, "host and port are required")
	}

	encodedUser := url.QueryEscape(user)
	encodedPassword := url.QueryEscape(password)

	uri := strings.Builder{}
	if encodedUser == "" || encodedPassword == "" {
		return "", NewErrorMongo(nil, "user and password are required")
	}

	uri.WriteString("mongodb://")
	uri.WriteString(encodedUser)
	uri.WriteString(":")
	uri.WriteString(encodedPassword)
	uri.WriteString("@")
	uri.WriteString(host)
	uri.WriteString(":")
	uri.WriteString(port)
	uri.WriteString("/?authSource=admin")

	return uri.String(), nil
}
