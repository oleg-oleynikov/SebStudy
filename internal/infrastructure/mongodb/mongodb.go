package mongodb

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	connectTimeout = time.Second * 15
)

type Config struct {
	URI      string
	Username string
	Password string
	Db       string
}

func NewMongoDbConn(ctx context.Context, cfg Config) (*mongo.Client, error) {

	client, err := mongo.Connect(context.Background(),
		options.Client().ApplyURI(cfg.URI).
			SetAuth(options.Credential{Username: cfg.Username, Password: cfg.Password}).
			SetConnectTimeout(connectTimeout))
	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	return client, err
}
