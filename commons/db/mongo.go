package db

import (
	"Expense_Manager/commons/config"
	"context"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func NewMongoDatabase() (*mongo.Database, error) {
	cfg := new(config.MongoDBConfig)
	err := config.GetConfig(cfg)
	if err != nil {
		return nil, errors.Wrap(err, "could not get mongodb config")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.ConnectionURI))
	if err != nil {
		return nil, errors.Wrap(err, "could not connect to mongodb")
	}

	return client.Database(cfg.DatabaseName), nil
}
